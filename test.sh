# queue
go test -run CreateQueue
go test -run ListQueue
go test -run GetQueueAttributes
go test -run SetQueueAttributes
go test -run CreateQueue
go test -run DeleteQueue

# queue message
go test -run SendMessage
go test -run BatchSendMessage
go test -run ReceiveMessage
go test -run BatchReceiveMessage
go test -run SendMessage
go test -run DeleteMessage
go test -run BatchSendMessage
go test -run BatchDeleteMessage

# topic
go test -run CreateTopic
go test -run ListTopic
go test -run GetTopicAttributes
go test -run SetTopicAttributes
go test -run CreateTopic
go test -run DeleteTopic

# subscription
go test -run Subscribe
go test -run ListSubscriptionByTopic
go test -run SetSubscriptionAttributes
go test -run GetSubscriptionAttributes
go test -run Subscribe
go test -run Unsubscribe

# topic message
go test -run PublishMessage
go test -run BatchPublishMessage
