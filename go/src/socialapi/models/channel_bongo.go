package models

import (
	"fmt"
	"time"

	"github.com/koding/bongo"

	"socialapi/request"
)

const ChannelTableName = "api.channel"

func (c Channel) GetId() int64 {
	return c.Id
}

func (c Channel) TableName() string {
	return ChannelTableName
}

func (c *Channel) AfterCreate() {
	bongo.B.AfterCreate(c)
}

func (c *Channel) AfterUpdate() {
	bongo.B.AfterUpdate(c)
}

func (c Channel) AfterDelete() {
	bongo.B.AfterDelete(c)
}

func (c *Channel) BeforeCreate() error {
	c.CreatedAt = time.Now().UTC()
	c.UpdatedAt = time.Now().UTC()
	c.DeletedAt = ZeroDate()
	c.Token = NewToken(c.CreatedAt).String()

	return c.MarkIfExempt()
}

func (c *Channel) BeforeUpdate() error {
	c.UpdatedAt = time.Now()

	return c.MarkIfExempt()
}

func (c *Channel) Update() error {
	if c.Name == "" || c.GroupName == "" {
		return fmt.Errorf("Validation failed Name: %s - GroupName:%s", c.Name, c.GroupName)
	}

	return bongo.B.Update(c)
}

func (c *Channel) Delete() error {
	if err := c.deleteChannelMessages(); err != nil {
		return err
	}

	if err := c.deleteChannelList(); err != nil {
		return err
	}

	return bongo.B.Delete(c)
}

func (c *Channel) ById(id int64) error {
	return bongo.B.ById(c, id)
}

func (c *Channel) One(q *bongo.Query) error {
	return bongo.B.One(c, c, q)
}

func (c *Channel) Some(data interface{}, q *bongo.Query) error {
	return bongo.B.Some(c, data, q)
}

func (c *Channel) FetchByIds(ids []int64) ([]Channel, error) {
	var channels []Channel

	if len(ids) == 0 {
		return channels, nil
	}

	if err := bongo.B.FetchByIds(c, &channels, ids); err != nil {
		return nil, err
	}
	return channels, nil
}

func (c *Channel) CountWithQuery(q *bongo.Query) (int, error) {
	return bongo.B.CountWithQuery(c, q)
}

func getMessageBatch(channelId int64, c int) ([]ChannelMessage, error) {
	messageIds, err := (&ChannelMessageList{}).
		FetchMessageIdsByChannelId(channelId, &request.Query{
		Skip:  c * 100,
		Limit: 100,
	})
	if err != nil {
		return nil, err
	}
	return (&ChannelMessage{}).FetchByIds(messageIds)
}

func isMessageCrossIndexed(messageId int64) (error, bool) {
	count, err := (&ChannelMessageList{}).CountWithQuery(&bongo.Query{
		Selector: map[string]interface{}{
			"message_id": messageId,
		},
	})
	if err != nil {
		return err, false
	}
	return nil, count > 1
}

func (c *Channel) deleteChannelMessages() error {
	for i := 0; ; i++ {
		messages, err := getMessageBatch(c.Id, i)
		if err != nil {
			return err
		}
		for _, message := range messages {
			err, isCrossIndexed := isMessageCrossIndexed(message.Id)
			if err != nil {
				return err
			}

			if isCrossIndexed {
				continue
			}

			if err = message.DeleteSilent(); err != nil {
				return err
			}
		}
		if len(messages) < 100 {
			return nil
		}
	}
}

func (c *Channel) deleteChannelList() error {
	cml := &ChannelMessageList{}
	err := cml.One(&bongo.Query{
		Selector: map[string]interface{}{"channel_id": c.Id}})

	if err != nil {
		return err
	}

	return cml.DeleteSilent()
}
