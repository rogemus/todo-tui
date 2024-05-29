# TODO
- [x] Add task
- [x] Remove task
- [x] Rename task
- [x] Mark as done
- [x] Show key bindings
- [ ] Save data to db
- [ ] Show Details
- [ ] Ability to add tags
- [ ] Tags autocomplete
- [ ] Edit on DetailsView
- [ ] Config in file
- [ ] Change keybindings in config
- [ ] Remote db

# Architecture

```
                    ┌───────────────┐                    
                    │   MainView    │                    
                    │ (tui_view.go) │                    
                    └───────┬───────┘                    
                            │                            
           ┌────────────────┴────────────────┐           
           │                                 │           
 ┌─────────┴──────────┐           ┌──────────┴─────────┐ 
 │     ListsView      │           │    SidebarView     │ 
 │  (lists_view.go)   │           │ (sidebar_view.go)  │ 
 └────────────────────┘           └────────────────────┘


 MainView                                                
┌──────────────────────────────┬───────────────────────┐ 
│ ListView                     │ DetailsView           │ 
│                              │                       │ 
│                              │                       │ 
│                              │                       │ 
│                              │                       │ 
└──────────────────────────────┴───────────────────────┘ 
```

# Mockup
```
┌───────────────────────────────────────────────────────────────────┬──────────────────────────────────────────┐
│                                                                   │                                          │
│    In Progress                                                    │  Title                                   │
│    ────────────────────────────────                               │  ──────────────────────────────────────  │
│   ┌─┐ Setup PostgreSQL for production dev                         │  #tag5 #tag2                             │
│   └─┘ #tag2                                                       │                                          │
│                                                                   │                                          │
│                                                                   │  Lorem ipsum dolor sit                   │
│ │  Todo                                                           │  amet, consectetur adipiscing elit.      │
│ │  ────────────────────────────────                               │  Donec a mauris rhoncus nunc vehicula    │
│ │ ┌─┐ Setup PostgreSQL for production server                      │  faucibus non auctor neque.              │
│ │ └─┘ #tag1                                                       │                                          │
│ │                                                                 │    * Lorem ipsum                         │
│ │                                                                 │    * Lorem ipsum 1                       │
│ │ ┌─┐ Setup PostgreSQL for production server                      │                                          │
│ │ └─┘ #tag1 #tag2 #tag3                                           │  Quisque eget lacus a ex sodales         │
│ │                                                                 │  accumsan. Quisque at sagittis ipsum.    │
│ │                                                                 │  Morbi consequat non est quis aliquam.   │
│ │                                                                 │  Morbi ac nisl sed lacus varius aliquet  │
│ │                              █ █ █                              │  sit amet vitae felis. Aenean vitae      │
│                                                                   │  nunc ut ligula fringilla rutrum.        │
│                                                                   │  Praesent rhoncus, ligula eget iaculis   │
│                                                                   │  accumsan, turpis odio viverra orci, a   │
│    Done                                                           │  faucibus nisl risus non nunc. Integer   │
│    ────────────────────────────────                               │  rutrum lorem nec ex gravida bibendum.   │
│   ╔═╗ Setup PostgreSQL for production server                      │                                          │
│   ╚═╝                                                             │                                          │
│        #tag5                                                      │                                          │
│                                                                   │                                          │
│                                                                   │                                          │
└───────────────────────────────────────────────────────────────────┴──────────────────────────────────────────┘
```

