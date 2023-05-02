# Go Mautic

The unofficial Go client for the [Mautic](https://www.mautic.org/) Rest API.

#### This is WIP project, we currently only support the endpoints for the following entities:
- [x] Webhooks
- [ ] Contacts (WIP)
- [ ] Campaigns (WIP)


Installation:
```
go get github.com/thiagoferolla/go-mautic
```

Usage:
```go
package main

import (
	"context"
	mautic "github.com/thiagoferolla/go-mautic"
)

func main() {
	ctx := context.Background()
	
	config := mautic.Config().
		SetBaseURL("https://your-mautic-url.com").
		SetUser("your-mautic-user").
		SetPassword("your-mautic-password")
	
	m, err := mautic.New(config)
	
	if err != nil {
		panic(err)
    }
	
	contacts, err := m.ListContacts(ctx)
	
	if err != nil {
        panic(err)
    }
	
	for _, contact := range contacts {
		fmt.Println(contact)
	}
}
```

Roadmap:
- [ ] Assets
- [ ] Categories
- [ ] Companies
- [ ] Dynamic Content
- [ ] Emails
- [ ] Fields
- [ ] Files
- [ ] Forms
- [ ] Marketing Messages
- [ ] Notes
- [ ] Notifications
- [ ] Pages
- [ ] Point Actions
- [ ] Point Triggers
- [ ] Reports
- [ ] Roles
- [ ] Segments
- [ ] Text Messages
- [ ] Stages
- [ ] Stats
- [ ] Tags
- [ ] Themes
- [ ] Tweets
- [ ] Users