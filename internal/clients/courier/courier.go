package courier

import (
	"context"
	"fmt"
	courierc "github.com/trycourier/courier-go/v3"
	courierclient "github.com/trycourier/courier-go/v3/client"
	option "github.com/trycourier/courier-go/v3/option"
	"notify/internal/config"
	e "notify/internal/entity"
)

type Client struct {
	client *courierclient.Client
}

func NewClient(cfg config.CourierConfig) *Client {
	c := courierclient.NewClient(option.WithAuthorizationToken(cfg.ApiKey))
	return &Client{client: c}
}

func sendMailRequestFromEmailMsg(msg e.EmailMsg) courierc.SendMessageRequest {
	return courierc.SendMessageRequest{
		Message: courierc.NewMessageFromContentMessage(&courierc.ContentMessage{
			To: courierc.NewMessageRecipientFromRecipient(
				courierc.NewRecipientFromUserRecipient(&courierc.UserRecipient{
					Email: courierc.String(msg.ToEmail),
				})),
			Content: courierc.NewContentFromElementalContentSugar(
				&courierc.ElementalContentSugar{
					Title: msg.Title,
					Body:  msg.Text,
				}),
		}),
	}
}

func (c *Client) SendTextMessage(ctx context.Context, msg e.EmailMsg) (*courierc.SendMessageResponse, error) {
	req := sendMailRequestFromEmailMsg(msg)
	resp, err := c.client.Send(ctx, &req)
	if err != nil {
		err = fmt.Errorf("%w :%w", ErrSendEmail, err)
	}
	return resp, err
}
