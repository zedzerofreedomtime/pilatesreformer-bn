# Pilates Reformer API

Go Gin API scaffold built from the frontend requirements in the public repository `zedzerofreedomtime/pilatesreformer`.

## Stack

- Go + Gin
- PostgreSQL
- Redis

## What this API covers

- public catalog bootstrap for the frontend
- role-based auth
- trainer application and admin approval flow
- trainer-scoped client listing
- admin CRUD for trainers, equipment, and home content
- booking quote calculation
- booking inquiry intake for admin follow-up

## Setup

1. Copy `.env.example` to `.env`
2. Start PostgreSQL and Redis
3. Apply `migrations/001_init.sql`
4. Apply `migrations/002_seed.sql`
5. Run `go run ./cmd/api`

## Demo credentials

- admin: `admin@reformrental.com` / `password123`
- user: `user@reformrental.com` / `password123`
- trainer: `trainer@reformrental.com` / `password123`

## Main endpoints

- `GET /api/v1/catalog/bootstrap`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `POST /api/v1/bookings/quote`
- `POST /api/v1/bookings/inquiries`
- `GET /api/v1/trainer/clients`
- `GET /api/v1/admin/trainer-applications`

See [docs/frontend-requirements.md](/E:/pilatesreformer-bn/docs/frontend-requirements.md) for the frontend-driven contract summary.
