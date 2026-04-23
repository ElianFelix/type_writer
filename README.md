# Type Writer

Elevate your typing mastery with our clean and sleek web app, designed to transform
your speed and accuracy into a professional-grade skill set. Whether you're looking
to test your limits or engage in targeted practice, our platform adapts to your
favorite texts, ensuring every session feels personal and engaging. With a diverse
range of activities specifically engineered to sharpen finger dexterity and precision,
you can seamlessly track your progress through an intuitive UI that stays out of your
way while you work. Stop just typing and start performing—experience a smarter,
more adaptive way to track your growth today.

## UI Screenshots

![Home Page](./resources/ui_screenshot_index.png)

## Technology Stack

### FrontEnd

- js
- vue
- vuetify
- pinia
- sass

### BackEnd

- go
- echo
- gorm
- postgresql

## How To Run

### FrontEnd

**requires**:
- node package manager -- this project uses pnpm

```bash
pnpm add
pnpm dev
navigate to http://localhost:3000
```

### BackEnd

**requires**:
- go >= 1.24,
- make 
- docker + docker-compose

```bash
go mod tidy
make deps-install
make local
```

## Dependencies

### FrontEnd

- js
- vue
- vuerouter
- vuetify
- pinia
- vite
- eslint
- sass
- axios

### BackEnd

- go
- go-sqlmock
- migrate
- echo
- echo-jwt
- slog-echo
- crypto
- gorm
- venom
- postgresql
- casbin
- crypto
- make
- wget
- docker
- docker compose

## Licensing

This repository is licensed under the Apache 2.0 license, see [LICENSE.txt](./LICENSE.txt) for more details.
