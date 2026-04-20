# Frontend Requirements Observed From GitHub Repo

Repository inspected: `https://github.com/zedzerofreedomtime/pilatesreformer`

## What the frontend currently needs

1. Public catalog data for equipment, trainers, rental plans, trainer service plans, and editable home page content
2. Role-based auth for `user`, `trainer`, and `admin`
3. User registration that can auto-login after success
4. Trainer registration that becomes a pending application for admin approval
5. Trainer dashboard data scoped to the logged-in trainer only
6. Admin CRUD for trainers, equipment, and home page content
7. Booking quote calculation that matches the frontend pricing rules:
   - installation fee is free from 3 months up
   - trainer-required machines force bundle sessions from `requiredSessions`
   - trainer-only pricing uses trainer session rate times selected trainer service plan
8. A natural follow-up endpoint for sending booking inquiries to admin or sales

## API shape implemented in this workspace

- `GET /healthz`
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `GET /api/v1/auth/me`
- `POST /api/v1/auth/logout`
- `GET /api/v1/catalog/bootstrap`
- `GET /api/v1/catalog/equipment`
- `GET /api/v1/catalog/trainers`
- `GET /api/v1/catalog/rental-plans`
- `GET /api/v1/catalog/trainer-service-plans`
- `GET /api/v1/catalog/home-content`
- `POST /api/v1/bookings/quote`
- `POST /api/v1/bookings/inquiries`
- `GET /api/v1/trainer/clients`
- `GET /api/v1/admin/trainer-applications`
- `POST /api/v1/admin/trainer-applications/:id/approve`
- `POST /api/v1/admin/trainer-applications/:id/reject`
- `POST /api/v1/admin/trainers`
- `PUT /api/v1/admin/trainers/:id`
- `DELETE /api/v1/admin/trainers/:id`
- `POST /api/v1/admin/equipment`
- `PUT /api/v1/admin/equipment/:id`
- `DELETE /api/v1/admin/equipment/:id`
- `GET /api/v1/admin/home-content`
- `PUT /api/v1/admin/home-content`
- `GET /api/v1/admin/bookings`
