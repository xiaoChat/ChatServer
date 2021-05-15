# api

## 用户相关

### Login 登录接口

登录接口

- Path: /login
- Method: `POST`
- Header: `无`
- Request
    | 名       | 类型   | 是否可传 | 描述 |
    | -------- | ------ | -------- | ---- |
    | username | string | 必填     |      |
    | password | string | 必填     |      |

- Response
    | 名            | 类型   | 描述     |
    | ------------- | ------ | -------- |
    | code          | int    | 响应状态 |
    | msg           | string | 响应消息 |
    | data          | Object | 响应数据 |
    | data.token    | string | 登录信息 |
    | data.userinfo | object | 用户信息 |
 
### Register 注册接口

注册接口

- Path: /register
- Method: `POST`
- Header: `无`
- Request
    | 名       | 类型   | 是否可传 | 描述 |
    | -------- | ------ | -------- | ---- |
    | email    | string | 必填     |      |
    | username | string | 必填     |      |
    | password | string | 必填     |      |

- Response
    | 名            | 类型   | 描述     |
    | ------------- | ------ | -------- |
    | code          | int    | 响应状态 |
    | msg           | string | 响应消息 |
    | data          | Object | 响应数据 |
    | data.token    | string | 登录信息 |
    | data.userinfo | object | 用户信息 |

## 资料

### getMyProfile 获取个人资料

获取个人资料

- Path: /user/getMyProfile
- Method: `POST`
- Header: `"{token: $TOKEN}"`
- Request: `无`

- Response
    | 名            | 类型   | 描述     |
    | ------------- | ------ | -------- |
    | code          | int    | 响应状态 |
    | msg           | string | 响应消息 |
    | data          | Object | 响应数据 |
    | data.userinfo | object | 用户信息 |


### getUserProfile 获取用户资料资料

获取用户资料资料

- Path: /user/getUserProfile
- Method: `POST`
- Header: `"{token: $TOKEN}"`
- Request
    | 名      | 类型  | 是否可传 | 描述                      |
    | ------- | ----- | -------- | ------------------------- |
    | user_id | array | 必填     | 用户的id，示例：[id1,id2] |

- Response
    | 名               | 类型   | 描述     |
    | ---------------- | ------ | -------- |
    | code             | int    | 响应状态 |
    | msg              | string | 响应消息 |
    | data             | array  | 响应数据 |
    | data[x].userinfo | object | 用户信息 |

### updateMyProfile 更新用户资料资料

更新用户资料资料

- Path: /user/updateMyProfile
- Method: `POST`
- Header: `"{token: $TOKEN}"`
- Request
    | 名                 | 类型   | 是否可传 | 描述                                               |
    | ------------------ | ------ | -------- | -------------------------------------------------- |
    | nick               | string | 可选     | 昵称                                               |
    | avatar             | string | 可选     | 头像地址                                           |
    | gender             | string | 可选     | 性别                                               |
    | selfSignature      | string | 可选     | 个性签名                                           |
    | allowType          | string | 可选     | 当被加人加好友时: 允许直接加为好友, 需要验证, 拒绝 |
    | birthday           | string | 可选     | 生日                                               |
    | location           | string | 可选     | 所在地                                             |
    | language           | string | 可选     | 语言                                               |
    | messageSettings    | string | 可选     | 消息设置                                           |
    | adminForbidType    | string | 可选     | 管理员禁止加好友标识                               |
    | level              | string | 可选     | 等级                                               |
    | role               | string | 可选     | 角色                                               |
    | profileCustomField | object | 可选     | 其他信息                                           |

- Response
    | 名   | 类型   | 描述     |
    | ---- | ------ | -------- |
    | code | int    | 响应状态 |
    | msg  | string | 响应消息 |
    | data | array  | 响应数据 |

### getBlacklist
### addToBlacklist
### removeFromBlacklist

## 消息

### createMessage 创建消息

- Path: /message/createMessage
- Method: `POST`
- Header: `"{token: $TOKEN}"`
- Request
  - createTextAtMessage 创建文本消息
    | 名               | 类型   | 是否可传 | 描述                           |
    | ---------------- | ------ | -------- | ------------------------------ |
    | to               | string | 必填     | 消息接收方的 userID 或 groupID |
    | conversationType | string | 必填     | 会话类型                       |
    | priority         | string | 必填     | 消息优先级                     |
    | payload          | object | 必填     | 消息内容                       |
    | payload.text     | text   | 必填     | 文本内容                       |
  - createTextAtMessage 创建带@提醒的消息
    | 名                 | 类型   | 是否可传 | 描述                           |
    | ------------------ | ------ | -------- | ------------------------------ |
    | to                 | string | 必填     | 消息接收方的 userID 或 groupID |
    | conversationType   | string | 必填     | 会话类型                       |
    | priority           | string | 必填     | 消息优先级                     |
    | payload            | object | 必填     | 消息内容                       |
    | payload.text       | text   | 必填     | 文本内容                       |
    | payload.atUserList | array  | 必填     | @提醒的用户 userID             |
  - createImageMessage 创建图片消息接口
  - createAudioMessage 创建音频消息接口
  - createVideoMessage 创建视频消息接口
  - createFileMessage 创建文件消息接口
    | 名               | 类型   | 是否可传 | 描述                           |
    | ---------------- | ------ | -------- | ------------------------------ |
    | to               | string | 必填     | 消息接收方的 userID 或 groupID |
    | conversationType | string | 必填     | 会话类型                       |
    | priority         | string | 必填     | 消息优先级                     |
    | payload          | object | 必填     | 消息内容                       |
    | payload.file     | object | 必填     | 文件                           |
  - createCustomMessage 创建自定义消息接口
    | 名                  | 类型   | 是否可传 | 描述                           |
    | ------------------- | ------ | -------- | ------------------------------ |
    | to                  | string | 必填     | 消息接收方的 userID 或 groupID |
    | conversationType    | string | 必填     | 会话类型                       |
    | priority            | string | 必填     | 消息优先级                     |
    | payload             | object | 必填     | 消息内容                       |
    | payload.data        | string | 必填     | 消息内容                       |
    | payload.description | string | 必填     | 描述                           |
    | payload.extension   | string | 必填     | 扩展                           |
  - createForwardMessage 创建转发消息
    | 名               | 类型   | 是否可传 | 描述                           |
    | ---------------- | ------ | -------- | ------------------------------ |
    | to               | string | 必填     | 消息接收方的 userID 或 groupID |
    | conversationType | string | 必填     | 会话类型                       |
    | priority         | string | 必填     | 消息优先级                     |
    | payload          | object | 必填     | 消息内容                       |
    | payload.message  | object | 必填     | 转发的消息内容                 |
  - createFaceMessage 创建表情包消息接口
    | 名               | 类型   | 是否可传 | 描述                           |
    | ---------------- | ------ | -------- | ------------------------------ |
    | to               | string | 必填     | 消息接收方的 userID 或 groupID |
    | conversationType | string | 必填     | 会话类型                       |
    | priority         | string | 必填     | 消息优先级                     |
    | payload          | object | 必填     | 消息内容                       |
    | payload.data     | string | 必填     | 消息内容                       |
    | payload.index    | string | 必填     | 索引                           |
  - createCustomMessage 创建自定义消息接口
    | 名                  | 类型   | 是否可传 | 描述                           |
    | ------------------- | ------ | -------- | ------------------------------ |
    | to                  | string | 必填     | 消息接收方的 userID 或 groupID |
    | conversationType    | string | 必填     | 会话类型                       |
    | priority            | string | 必填     | 消息优先级                     |
    | payload             | object | 必填     | 消息内容                       |
    | payload.data        | string | 必填     | 消息内容                       |
    | payload.description | string | 必填     | 描述                           |
    | payload.extension   | string | 必填     | 扩展                           |

- Response
    | 名   | 类型   | 描述                                   |
    | ---- | ------ | -------------------------------------- |
    | code | int    | 响应状态                               |
    | msg  | string | 响应消息                               |
    | data | Object | 响应数据, 具体查看附录 `附录1 MESSAGE` |

### sendMessage
### revokeMessage
### resendMessage
### getMessageList 获取消息列表

- Path: /message/getMessageList
- Method: `POST`
- Header: `"{token: $TOKEN}"`
- Request
    | 名             | 类型   | 是否可传 | 描述   |
    | -------------- | ------ | -------- | ------ |
    | conversationID | string | 必填     | 会话id |
    | count          | int    | 可选     | 条数   |

- Response
    | 名              | 类型   | 描述                                   |
    | --------------- | ------ | -------------------------------------- |
    | code            | int    | 响应状态                               |
    | msg             | string | 响应消息                               |
    | data            | array  | 响应数据                               |
    | data[x].message | Object | 响应数据, 具体查看附录 `附录1 MESSAGE` |

### setMessageRead

## 会话
### getConversationList 获取会话列表

- Path: /conversation/getConversationList
- Method: `POST`
- Header: `"{token: $TOKEN}"`
- Request
    | 名             | 类型   | 是否可传 | 描述   |
    | -------------- | ------ | -------- | ------ |
    | conversationID | string | 必填     | 会话id |
    | count          | int    | 可选     | 条数   |

- Response
    | 名              | 类型   | 描述                                   |
    | --------------- | ------ | -------------------------------------- |
    | code            | int    | 响应状态                               |
    | msg             | string | 响应消息                               |
    | data            | array  | 响应数据                               |
    | data[x].message | Object | 响应数据, 具体查看附录 `附录1 MESSAGE` |

### getConversationProfile
### deleteConversation
## 群组
### getGroupList
### getGroupProfile
### createGroup
### dismissGroup
### updateGroupProfile
### joinGroup
### quitGroup
### searchGroupByID
### changeGroupOwner
### handleGroupApplication
### setMessageRemindType
## 群成员
### getGroupMemberList
### getGroupMemberProfile
### getGroupOnlineMemberCount
### addGroupMember
### deleteGroupMember
### setGroupMemberMuteTime
### setGroupMemberRole
### setGroupMemberNameCard
### setGroupMemberCustomField

## 附录

### 附录1 MESSAGE

| 名               | 类型    | 描述                                                       |
| ---------------- | ------- | ---------------------------------------------------------- |
| ID               | string  | 消息id                                                     |
| type             | string  | 会话类型类型：                                             |
| payload          | object  | 消息内容                                                   |
| conversationID   | string  | 会话类型                                                   |
| conversationType | string  | 消息优先级                                                 |
| to               | string  | 接收方的 userID                                            |
| from             | string  | 发送方的 userID，在消息发送时，会默认设置为当前登录的用户  |
| flow             | string  | 消息的流向 in 为收到的消息 out 为发出的消息                |
| time             | number  | 消息时间戳。单位：秒                                       |
| status           | string  | 消息状态。 unSend(未发送) success(发送成功) fail(发送失败) |
| isRevoked        | boolean | false 是否被撤回的消息，true 标识被撤回的消息              |
| priority         | string  | 消息优先级                                                 |
| nick             | string  | 消息发送者的昵称                                           |
| avatar           | string  | 消息发送者的头像地址                                       |
| isPeerRead       | boolean | false C2C 消息对端是否已读，true 标识对端已读              |
| nameCard         | string  | 消息发送者的群名片                                         |
| atUserList       | array   | 群聊时此字段存储被 at 的群成员的 userID                    |
| cloudCustomData  | string  | 消息自定义数据                                             |

### 附录2 一些数据类型
- conversationType
- TIM.TYPES.CONV_C2C	C2C(Client to Client, 端到端) 会话
TIM.TYPES.CONV_GROUP	GROUP(群组) 会话
TIM.TYPES.CONV_SYSTEM	SYSTEM(系统) 会话