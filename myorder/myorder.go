package myorder

import (
	"errors"
	"fmt"
	"github.com/chilts/sid"
	"github.com/kjk/betterguid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"math/rand"
	"sync"
	"time"
)

// To run:
// go run main.go

func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:           %s\n", id.String())
}

func genKsuid() {
	id := ksuid.New()
	fmt.Printf("OrderIdTest github.com/segmentio/ksuid:  %s\n", id.String())
}

func genBetterGUID() {
	id := betterguid.New()
	fmt.Printf("OrderIdTest github.com/kjk/betterguid:   %s\n", id)
}

func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("OrderIdTest github.com/oklog/ulid:       %s\n", id.String())
}

func genSonyflake() {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		fmt.Errorf("OrderIdTest flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	fmt.Printf("OrderIdTest github.com/sony/sonyflake:   %x\n", id)
}

func genSid() {
	id := sid.Id()
	fmt.Printf("OrderIdTest github.com/chilts/sid:       %s\n", id)
}

func genUUIDv4() {
	id := uuid.NewV4()
	fmt.Printf("OrderIdTest github.com/satori/go.uuid:   %s\n", id)
}

func GetOrderIdTest() {
	genXid()
	genKsuid()
	genBetterGUID()
	genUlid()
	genSonyflake()
	genSid()
	genUUIDv4()
}

func GenSnowFlake() {
	currWoker := &IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	newId, _ := currWoker.NextId()
	fmt.Printf("OrderIdTest GenSnowFlake %d", newId)
}

type IdWorker struct {
	startTime             int64
	workerIdBits          uint
	datacenterIdBits      uint
	maxWorkerId           int64
	maxDatacenterId       int64
	sequenceBits          uint
	workerIdLeftShift     uint
	datacenterIdLeftShift uint
	timestampLeftShift    uint
	sequenceMask          int64
	workerId              int64
	datacenterId          int64
	sequence              int64
	lastTimestamp         int64
	signMask              int64
	idLock                *sync.Mutex
}

func (this *IdWorker) InitIdWorker(workerId, datacenterId int64) error {

	var baseValue int64 = -1
	this.startTime = 1463834116272
	this.workerIdBits = 5
	this.datacenterIdBits = 5
	this.maxWorkerId = baseValue ^ (baseValue << this.workerIdBits)
	this.maxDatacenterId = baseValue ^ (baseValue << this.datacenterIdBits)
	this.sequenceBits = 12
	this.workerIdLeftShift = this.sequenceBits
	this.datacenterIdLeftShift = this.workerIdBits + this.workerIdLeftShift
	this.timestampLeftShift = this.datacenterIdBits + this.datacenterIdLeftShift
	this.sequenceMask = baseValue ^ (baseValue << this.sequenceBits)
	this.sequence = 0
	this.lastTimestamp = -1
	this.signMask = ^baseValue + 1

	this.idLock = &sync.Mutex{}

	if this.workerId < 0 || this.workerId > this.maxWorkerId {
		return errors.New(fmt.Sprintf("workerId[%v] is less than 0 or greater than maxWorkerId[%v].", workerId, datacenterId))
	}
	if this.datacenterId < 0 || this.datacenterId > this.maxDatacenterId {
		return errors.New(fmt.Sprintf("datacenterId[%d] is less than 0 or greater than maxDatacenterId[%d].", workerId, datacenterId))
	}
	this.workerId = workerId
	this.datacenterId = datacenterId
	return nil
}

func (this *IdWorker) NextId() (int64, error) {
	this.idLock.Lock()
	timestamp := time.Now().UnixNano()
	if timestamp < this.lastTimestamp {
		return -1, errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", this.lastTimestamp-timestamp))
	}

	if timestamp == this.lastTimestamp {
		this.sequence = (this.sequence + 1) & this.sequenceMask
		if this.sequence == 0 {
			timestamp = this.tilNextMillis()
			this.sequence = 0
		}
	} else {
		this.sequence = 0
	}

	this.lastTimestamp = timestamp

	this.idLock.Unlock()

	id := ((timestamp - this.startTime) << this.timestampLeftShift) |
		(this.datacenterId << this.datacenterIdLeftShift) |
		(this.workerId << this.workerIdLeftShift) |
		this.sequence

	if id < 0 {
		id = -id
	}

	return id, nil
}

func (this *IdWorker) tilNextMillis() int64 {
	timestamp := time.Now().UnixNano()
	if timestamp <= this.lastTimestamp {
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return timestamp
}
