package im

const (
	DefaultRole    = 0 //默认
	Lecturer       = 1 //主讲
	TeacherOfClass = 2 //班主任
	Encourager     = 3 //鼓励师

	Student = 4 //学生
)

//字体类
type (
	ChatNode struct {
		//聊天发送方
		MsgId    int64
		TmStamp  int64
		FromUid  int64
		FromRole int
		Content  string
		//回复某条记录
		ReplySwith   int
		ReployToUid  int64
		ReplyContent string

		//聊天附属属性
		AttachAttr ChatAttachAttr
	}
	ChatAttachAttr struct {
		FontAttr  string //字体类属性	学生端
		BadgeAttr string //徽章类		学生端
		TitleAttr string //课中称号类  	学生端
		VipAttr   string //续保 			主讲端
	}
	/*
			存储样式
			三层存储
			维度 lessonId_liveroomId
			主讲端：list(200条消息一链、分多链)
			班主任侧：classid维度 list(20条消息一链、分多链)
			学生端：nodeid维度 list(100条消息一链、分多链)

			存储控制部分
			维度 lessonId_liveroomId
			主讲端:key-value（cur_index当前链下标）
			班主任侧:hash field(classid)  value(cur_index当前班级的下标)
			学生端:hash field(nodeId) value(cur_index当前班级的下标)

			msgId问题

		   1.实现方案图
		   2.功能通信流程图
	*/

	// 禁言/解禁
	/*
		个体禁言
		主讲禁言:
		hash(控制部分：记录有哪些被禁言的学生的班级)
				key：lessonid_liveroomId  field：nodeId(班主任管理的节点)  value：时间戳
		zset(数据部分：记录每个班级被禁言的学生)
				key：lessonId_liveroomId_nodeId member：uid（被禁言的学员） score：时间戳
		班主任禁言：
		hash(控制部分：记录有哪些被禁言的学生的班级)
				key：lessonid_liveroomId  field：nodeId(班主任管理的节点)  value：时间戳
		zset(数据部分：记录每个班级被禁言的学生)
				key：lessonId_liveroomId_nodeId member：uid（被禁言的学员） score：时间戳


		团体禁言
		主讲禁言：
			hash(控制部分：记录有哪些被禁言的学生的班级)
					key：lessonid_liveroomId  field：nodeId(班主任管理的节点)  value：时间戳
		班主任禁言：
			hash(控制部分：记录有哪些被禁言的学生的班级)
					key：lessonid_liveroomId  field：nodeId(班主任管理的节点)  value：时间戳
	*/

)
