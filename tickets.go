// tickets.go
package main

import (
    "context"
    "errors"
)

type Ticket struct {
    ID      int64
    OwnerID int64
}

func GetTicketByID(ctx context.Context, id int64) (*Ticket, error) {
    return &Ticket{ID: id}, nil
}

func CanWriteTicket(ctx context.Context, t *Ticket) bool {
    return true // stub: replace with a real ownership check
}

func DeleteTicket(ctx context.Context, t *Ticket) error {
    return nil
}

func removeTicket(ctx context.Context, ticketID int64) error {
    ticket, err := GetTicketByID(ctx, ticketID)
    if err != nil {
        return err
    }
    allowed := CanWriteTicket(ctx, ticket)
    if !allowed {
        return errors.New("forbidden")
    }
    return DeleteTicket(ctx, ticket)
}
