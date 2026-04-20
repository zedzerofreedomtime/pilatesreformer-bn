create extension if not exists "pgcrypto";

create table if not exists users (
  id text primary key,
  name text not null,
  email text not null unique,
  phone text not null default '',
  role_id text not null check (role_id in ('user', 'trainer', 'admin')),
  password_hash text not null,
  trainer_id text,
  created_at timestamptz not null default now()
);

create table if not exists equipment (
  id text primary key,
  name text not null,
  image text not null,
  badge text not null,
  monthly_rate integer not null,
  trainer_mode text not null check (trainer_mode in ('optional', 'required')),
  summary text not null,
  ideal_for text not null,
  footprint text not null,
  features text[] not null default '{}',
  created_at timestamptz not null default now()
);

create table if not exists trainers (
  id text primary key,
  name text not null,
  image text not null,
  specialty text not null,
  session_rate integer not null,
  availability text not null,
  summary text not null,
  schedule_window text not null,
  available_slots integer not null default 0,
  booked_slots integer not null default 0,
  machine_focus text[] not null default '{}',
  exercise_focus text[] not null default '{}',
  weekly_schedule jsonb not null default '[]'::jsonb,
  created_at timestamptz not null default now()
);

create table if not exists trainer_clients (
  id text primary key,
  trainer_id text not null references trainers(id) on delete cascade,
  client_name text not null,
  equipment_name text not null,
  plan_name text not null,
  next_session text not null,
  contact text not null,
  status text not null,
  created_at timestamptz not null default now()
);

create table if not exists trainer_applications (
  id text primary key,
  name text not null,
  email text not null,
  phone text not null,
  password_hash text not null,
  specialty text not null,
  machine_focus text[] not null default '{}',
  status text not null default 'pending' check (status in ('pending', 'approved', 'rejected')),
  submitted_at timestamptz not null default now()
);

create table if not exists rental_plans (
  id text primary key,
  name text not null,
  months integer not null,
  discount numeric(5,2) not null,
  optional_sessions integer not null,
  required_sessions integer not null,
  note text not null
);

create table if not exists trainer_service_plans (
  id text primary key,
  name text not null,
  sessions integer not null,
  discount numeric(5,2) not null,
  note text not null
);

create table if not exists home_page_contents (
  id text primary key,
  payload jsonb not null,
  updated_at timestamptz not null default now()
);

create table if not exists booking_inquiries (
  id text primary key,
  name text not null,
  email text not null,
  phone text not null,
  mode_id text not null check (mode_id in ('bundle', 'equipment-only', 'trainer-only')),
  equipment_id text,
  rental_plan_id text,
  trainer_id text,
  trainer_service_plan_id text,
  grand_total integer not null,
  notes text not null default '',
  status text not null default 'new',
  created_at timestamptz not null default now()
);
