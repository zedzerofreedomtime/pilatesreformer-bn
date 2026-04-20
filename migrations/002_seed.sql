insert into users (id, name, email, phone, role_id, password_hash, trainer_id)
values
  ('user-admin', 'Admin', 'admin@reformrental.com', '0800000000', 'admin', crypt('password123', gen_salt('bf')), null),
  ('user-member-demo', 'Demo User', 'user@reformrental.com', '0810000000', 'user', crypt('password123', gen_salt('bf')), null),
  ('user-trainer-demo', 'Trainer Demo', 'trainer@reformrental.com', '0820000000', 'trainer', crypt('password123', gen_salt('bf')), 'coach-pim')
on conflict do nothing;

insert into equipment (id, name, image, badge, monthly_rate, trainer_mode, summary, ideal_for, footprint, features)
values
  ('reformer', 'Pilates Reformer', '/images/equipment-reformer.svg', 'Home friendly', 12900, 'optional', 'Popular home pilates machine with flexible usage', 'Beginners, posture work, home studio', '2.4m x 0.8m', '{"Easy for home onboarding","Works for core and flexibility","Trainer can be added later"}'),
  ('tower', 'Cadillac / Tower', '/images/equipment-tower.svg', 'Trainer required', 18500, 'required', 'Specialized machine for rehab and corrective movement', 'Private coaching and recovery programs', '2.8m x 1.1m', '{"Requires trainer guidance","Best for corrective movement","Setup support included"}'),
  ('chair', 'Stability Chair', '/images/equipment-chair.svg', 'Trainer required', 14900, 'required', 'Compact but technical machine for balance and strength', 'Small spaces and lower-body work', '1.6m x 1.4m', '{"High control work","Balance and strength focus"}'),
  ('functional', 'Functional Trainer', '/images/equipment-functional.svg', 'Trainer required', 16900, 'required', 'Cable-based machine for strength and mobility training', 'Performance and hybrid coaching', '2.2m x 1.8m', '{"Coach-guided loading","Great for athletic programs"}')
on conflict do nothing;

insert into trainers (id, name, image, specialty, session_rate, availability, summary, schedule_window, available_slots, booked_slots, machine_focus, exercise_focus, weekly_schedule)
values
  ('coach-pim', 'Coach Pim', '/images/trainer-pim.svg', 'Reformer foundation and posture reset', 1800, 'Onsite Bangkok / Online cueing', 'Strong fit for first-time reformer users who want a reliable foundation', 'Sunday - Saturday 08:00 - 17:00', 52, 11, '{"Pilates Reformer","Stability Chair"}', '{"Posture reset","Core foundation","Beginner pilates"}', '[{"id":"sun","label":"Sunday","shortLabel":"Sun","availableCount":7,"bookedCount":2,"slots":[{"key":"sun-08:00","label":"08:00 - 09:00","status":"available"},{"key":"sun-09:00","label":"09:00 - 10:00","status":"booked"},{"key":"sun-10:00","label":"10:00 - 11:00","status":"booked"},{"key":"sun-11:00","label":"11:00 - 12:00","status":"available"},{"key":"sun-12:00","label":"12:00 - 13:00","status":"available"},{"key":"sun-13:00","label":"13:00 - 14:00","status":"available"},{"key":"sun-14:00","label":"14:00 - 15:00","status":"available"},{"key":"sun-15:00","label":"15:00 - 16:00","status":"available"},{"key":"sun-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"mon","label":"Monday","shortLabel":"Mon","availableCount":7,"bookedCount":2,"slots":[{"key":"mon-08:00","label":"08:00 - 09:00","status":"available"},{"key":"mon-09:00","label":"09:00 - 10:00","status":"available"},{"key":"mon-10:00","label":"10:00 - 11:00","status":"available"},{"key":"mon-11:00","label":"11:00 - 12:00","status":"available"},{"key":"mon-12:00","label":"12:00 - 13:00","status":"available"},{"key":"mon-13:00","label":"13:00 - 14:00","status":"booked"},{"key":"mon-14:00","label":"14:00 - 15:00","status":"booked"},{"key":"mon-15:00","label":"15:00 - 16:00","status":"available"},{"key":"mon-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"tue","label":"Tuesday","shortLabel":"Tue","availableCount":8,"bookedCount":1,"slots":[{"key":"tue-08:00","label":"08:00 - 09:00","status":"booked"},{"key":"tue-09:00","label":"09:00 - 10:00","status":"available"},{"key":"tue-10:00","label":"10:00 - 11:00","status":"available"},{"key":"tue-11:00","label":"11:00 - 12:00","status":"available"},{"key":"tue-12:00","label":"12:00 - 13:00","status":"available"},{"key":"tue-13:00","label":"13:00 - 14:00","status":"available"},{"key":"tue-14:00","label":"14:00 - 15:00","status":"available"},{"key":"tue-15:00","label":"15:00 - 16:00","status":"available"},{"key":"tue-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"wed","label":"Wednesday","shortLabel":"Wed","availableCount":7,"bookedCount":2,"slots":[{"key":"wed-08:00","label":"08:00 - 09:00","status":"available"},{"key":"wed-09:00","label":"09:00 - 10:00","status":"available"},{"key":"wed-10:00","label":"10:00 - 11:00","status":"available"},{"key":"wed-11:00","label":"11:00 - 12:00","status":"booked"},{"key":"wed-12:00","label":"12:00 - 13:00","status":"booked"},{"key":"wed-13:00","label":"13:00 - 14:00","status":"available"},{"key":"wed-14:00","label":"14:00 - 15:00","status":"available"},{"key":"wed-15:00","label":"15:00 - 16:00","status":"available"},{"key":"wed-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"thu","label":"Thursday","shortLabel":"Thu","availableCount":8,"bookedCount":1,"slots":[{"key":"thu-08:00","label":"08:00 - 09:00","status":"available"},{"key":"thu-09:00","label":"09:00 - 10:00","status":"available"},{"key":"thu-10:00","label":"10:00 - 11:00","status":"available"},{"key":"thu-11:00","label":"11:00 - 12:00","status":"available"},{"key":"thu-12:00","label":"12:00 - 13:00","status":"available"},{"key":"thu-13:00","label":"13:00 - 14:00","status":"available"},{"key":"thu-14:00","label":"14:00 - 15:00","status":"available"},{"key":"thu-15:00","label":"15:00 - 16:00","status":"booked"},{"key":"thu-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"fri","label":"Friday","shortLabel":"Fri","availableCount":7,"bookedCount":2,"slots":[{"key":"fri-08:00","label":"08:00 - 09:00","status":"available"},{"key":"fri-09:00","label":"09:00 - 10:00","status":"booked"},{"key":"fri-10:00","label":"10:00 - 11:00","status":"booked"},{"key":"fri-11:00","label":"11:00 - 12:00","status":"available"},{"key":"fri-12:00","label":"12:00 - 13:00","status":"available"},{"key":"fri-13:00","label":"13:00 - 14:00","status":"available"},{"key":"fri-14:00","label":"14:00 - 15:00","status":"available"},{"key":"fri-15:00","label":"15:00 - 16:00","status":"available"},{"key":"fri-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"sat","label":"Saturday","shortLabel":"Sat","availableCount":8,"bookedCount":1,"slots":[{"key":"sat-08:00","label":"08:00 - 09:00","status":"available"},{"key":"sat-09:00","label":"09:00 - 10:00","status":"available"},{"key":"sat-10:00","label":"10:00 - 11:00","status":"available"},{"key":"sat-11:00","label":"11:00 - 12:00","status":"available"},{"key":"sat-12:00","label":"12:00 - 13:00","status":"available"},{"key":"sat-13:00","label":"13:00 - 14:00","status":"available"},{"key":"sat-14:00","label":"14:00 - 15:00","status":"available"},{"key":"sat-15:00","label":"15:00 - 16:00","status":"available"},{"key":"sat-16:00","label":"16:00 - 17:00","status":"booked"}]}]'),
  ('coach-tone', 'Coach Tone', '/images/trainer-tone.svg', 'Strength, mobility, and functional movement', 2200, 'Onsite Bangkok and nearby area', 'Well suited for strength-oriented and technical equipment programs', 'Sunday - Saturday 08:00 - 17:00', 51, 12, '{"Functional Trainer","Cadillac / Tower"}', '{"Strength training","Mobility","Functional movement"}', '[{"id":"sun","label":"Sunday","shortLabel":"Sun","availableCount":7,"bookedCount":2,"slots":[{"key":"sun-08:00","label":"08:00 - 09:00","status":"available"},{"key":"sun-09:00","label":"09:00 - 10:00","status":"available"},{"key":"sun-10:00","label":"10:00 - 11:00","status":"available"},{"key":"sun-11:00","label":"11:00 - 12:00","status":"available"},{"key":"sun-12:00","label":"12:00 - 13:00","status":"available"},{"key":"sun-13:00","label":"13:00 - 14:00","status":"available"},{"key":"sun-14:00","label":"14:00 - 15:00","status":"available"},{"key":"sun-15:00","label":"15:00 - 16:00","status":"booked"},{"key":"sun-16:00","label":"16:00 - 17:00","status":"booked"}]},{"id":"mon","label":"Monday","shortLabel":"Mon","availableCount":8,"bookedCount":1,"slots":[{"key":"mon-08:00","label":"08:00 - 09:00","status":"available"},{"key":"mon-09:00","label":"09:00 - 10:00","status":"available"},{"key":"mon-10:00","label":"10:00 - 11:00","status":"available"},{"key":"mon-11:00","label":"11:00 - 12:00","status":"available"},{"key":"mon-12:00","label":"12:00 - 13:00","status":"available"},{"key":"mon-13:00","label":"13:00 - 14:00","status":"available"},{"key":"mon-14:00","label":"14:00 - 15:00","status":"available"},{"key":"mon-15:00","label":"15:00 - 16:00","status":"available"},{"key":"mon-16:00","label":"16:00 - 17:00","status":"booked"}]},{"id":"tue","label":"Tuesday","shortLabel":"Tue","availableCount":7,"bookedCount":2,"slots":[{"key":"tue-08:00","label":"08:00 - 09:00","status":"available"},{"key":"tue-09:00","label":"09:00 - 10:00","status":"booked"},{"key":"tue-10:00","label":"10:00 - 11:00","status":"booked"},{"key":"tue-11:00","label":"11:00 - 12:00","status":"available"},{"key":"tue-12:00","label":"12:00 - 13:00","status":"available"},{"key":"tue-13:00","label":"13:00 - 14:00","status":"available"},{"key":"tue-14:00","label":"14:00 - 15:00","status":"available"},{"key":"tue-15:00","label":"15:00 - 16:00","status":"available"},{"key":"tue-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"wed","label":"Wednesday","shortLabel":"Wed","availableCount":8,"bookedCount":1,"slots":[{"key":"wed-08:00","label":"08:00 - 09:00","status":"available"},{"key":"wed-09:00","label":"09:00 - 10:00","status":"available"},{"key":"wed-10:00","label":"10:00 - 11:00","status":"available"},{"key":"wed-11:00","label":"11:00 - 12:00","status":"available"},{"key":"wed-12:00","label":"12:00 - 13:00","status":"available"},{"key":"wed-13:00","label":"13:00 - 14:00","status":"available"},{"key":"wed-14:00","label":"14:00 - 15:00","status":"booked"},{"key":"wed-15:00","label":"15:00 - 16:00","status":"available"},{"key":"wed-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"thu","label":"Thursday","shortLabel":"Thu","availableCount":7,"bookedCount":2,"slots":[{"key":"thu-08:00","label":"08:00 - 09:00","status":"booked"},{"key":"thu-09:00","label":"09:00 - 10:00","status":"booked"},{"key":"thu-10:00","label":"10:00 - 11:00","status":"available"},{"key":"thu-11:00","label":"11:00 - 12:00","status":"available"},{"key":"thu-12:00","label":"12:00 - 13:00","status":"available"},{"key":"thu-13:00","label":"13:00 - 14:00","status":"available"},{"key":"thu-14:00","label":"14:00 - 15:00","status":"available"},{"key":"thu-15:00","label":"15:00 - 16:00","status":"available"},{"key":"thu-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"fri","label":"Friday","shortLabel":"Fri","availableCount":8,"bookedCount":1,"slots":[{"key":"fri-08:00","label":"08:00 - 09:00","status":"available"},{"key":"fri-09:00","label":"09:00 - 10:00","status":"available"},{"key":"fri-10:00","label":"10:00 - 11:00","status":"available"},{"key":"fri-11:00","label":"11:00 - 12:00","status":"available"},{"key":"fri-12:00","label":"12:00 - 13:00","status":"available"},{"key":"fri-13:00","label":"13:00 - 14:00","status":"booked"},{"key":"fri-14:00","label":"14:00 - 15:00","status":"available"},{"key":"fri-15:00","label":"15:00 - 16:00","status":"available"},{"key":"fri-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"sat","label":"Saturday","shortLabel":"Sat","availableCount":6,"bookedCount":3,"slots":[{"key":"sat-08:00","label":"08:00 - 09:00","status":"available"},{"key":"sat-09:00","label":"09:00 - 10:00","status":"available"},{"key":"sat-10:00","label":"10:00 - 11:00","status":"booked"},{"key":"sat-11:00","label":"11:00 - 12:00","status":"booked"},{"key":"sat-12:00","label":"12:00 - 13:00","status":"booked"},{"key":"sat-13:00","label":"13:00 - 14:00","status":"available"},{"key":"sat-14:00","label":"14:00 - 15:00","status":"available"},{"key":"sat-15:00","label":"15:00 - 16:00","status":"available"},{"key":"sat-16:00","label":"16:00 - 17:00","status":"available"}]}]'),
  ('coach-fon', 'Coach Fon', '/images/trainer-fon.svg', 'Private rehab flow and breath-led pilates', 2400, 'Weekend onsite / Hybrid follow-up', 'Focused on deeper private sessions, alignment, and recovery-oriented movement', 'Sunday - Saturday 08:00 - 17:00', 53, 10, '{"Cadillac / Tower","Pilates Reformer"}', '{"Rehab flow","Breath-led pilates","Alignment work"}', '[{"id":"sun","label":"Sunday","shortLabel":"Sun","availableCount":8,"bookedCount":1,"slots":[{"key":"sun-08:00","label":"08:00 - 09:00","status":"available"},{"key":"sun-09:00","label":"09:00 - 10:00","status":"available"},{"key":"sun-10:00","label":"10:00 - 11:00","status":"available"},{"key":"sun-11:00","label":"11:00 - 12:00","status":"available"},{"key":"sun-12:00","label":"12:00 - 13:00","status":"available"},{"key":"sun-13:00","label":"13:00 - 14:00","status":"booked"},{"key":"sun-14:00","label":"14:00 - 15:00","status":"available"},{"key":"sun-15:00","label":"15:00 - 16:00","status":"available"},{"key":"sun-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"mon","label":"Monday","shortLabel":"Mon","availableCount":7,"bookedCount":2,"slots":[{"key":"mon-08:00","label":"08:00 - 09:00","status":"available"},{"key":"mon-09:00","label":"09:00 - 10:00","status":"available"},{"key":"mon-10:00","label":"10:00 - 11:00","status":"booked"},{"key":"mon-11:00","label":"11:00 - 12:00","status":"booked"},{"key":"mon-12:00","label":"12:00 - 13:00","status":"available"},{"key":"mon-13:00","label":"13:00 - 14:00","status":"available"},{"key":"mon-14:00","label":"14:00 - 15:00","status":"available"},{"key":"mon-15:00","label":"15:00 - 16:00","status":"available"},{"key":"mon-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"tue","label":"Tuesday","shortLabel":"Tue","availableCount":8,"bookedCount":1,"slots":[{"key":"tue-08:00","label":"08:00 - 09:00","status":"available"},{"key":"tue-09:00","label":"09:00 - 10:00","status":"available"},{"key":"tue-10:00","label":"10:00 - 11:00","status":"available"},{"key":"tue-11:00","label":"11:00 - 12:00","status":"available"},{"key":"tue-12:00","label":"12:00 - 13:00","status":"available"},{"key":"tue-13:00","label":"13:00 - 14:00","status":"available"},{"key":"tue-14:00","label":"14:00 - 15:00","status":"booked"},{"key":"tue-15:00","label":"15:00 - 16:00","status":"available"},{"key":"tue-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"wed","label":"Wednesday","shortLabel":"Wed","availableCount":7,"bookedCount":2,"slots":[{"key":"wed-08:00","label":"08:00 - 09:00","status":"available"},{"key":"wed-09:00","label":"09:00 - 10:00","status":"available"},{"key":"wed-10:00","label":"10:00 - 11:00","status":"available"},{"key":"wed-11:00","label":"11:00 - 12:00","status":"available"},{"key":"wed-12:00","label":"12:00 - 13:00","status":"available"},{"key":"wed-13:00","label":"13:00 - 14:00","status":"available"},{"key":"wed-14:00","label":"14:00 - 15:00","status":"available"},{"key":"wed-15:00","label":"15:00 - 16:00","status":"booked"},{"key":"wed-16:00","label":"16:00 - 17:00","status":"booked"}]},{"id":"thu","label":"Thursday","shortLabel":"Thu","availableCount":7,"bookedCount":2,"slots":[{"key":"thu-08:00","label":"08:00 - 09:00","status":"available"},{"key":"thu-09:00","label":"09:00 - 10:00","status":"available"},{"key":"thu-10:00","label":"10:00 - 11:00","status":"booked"},{"key":"thu-11:00","label":"11:00 - 12:00","status":"booked"},{"key":"thu-12:00","label":"12:00 - 13:00","status":"available"},{"key":"thu-13:00","label":"13:00 - 14:00","status":"available"},{"key":"thu-14:00","label":"14:00 - 15:00","status":"available"},{"key":"thu-15:00","label":"15:00 - 16:00","status":"available"},{"key":"thu-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"fri","label":"Friday","shortLabel":"Fri","availableCount":8,"bookedCount":1,"slots":[{"key":"fri-08:00","label":"08:00 - 09:00","status":"booked"},{"key":"fri-09:00","label":"09:00 - 10:00","status":"available"},{"key":"fri-10:00","label":"10:00 - 11:00","status":"available"},{"key":"fri-11:00","label":"11:00 - 12:00","status":"available"},{"key":"fri-12:00","label":"12:00 - 13:00","status":"available"},{"key":"fri-13:00","label":"13:00 - 14:00","status":"available"},{"key":"fri-14:00","label":"14:00 - 15:00","status":"available"},{"key":"fri-15:00","label":"15:00 - 16:00","status":"available"},{"key":"fri-16:00","label":"16:00 - 17:00","status":"available"}]},{"id":"sat","label":"Saturday","shortLabel":"Sat","availableCount":7,"bookedCount":2,"slots":[{"key":"sat-08:00","label":"08:00 - 09:00","status":"available"},{"key":"sat-09:00","label":"09:00 - 10:00","status":"booked"},{"key":"sat-10:00","label":"10:00 - 11:00","status":"available"},{"key":"sat-11:00","label":"11:00 - 12:00","status":"available"},{"key":"sat-12:00","label":"12:00 - 13:00","status":"available"},{"key":"sat-13:00","label":"13:00 - 14:00","status":"booked"},{"key":"sat-14:00","label":"14:00 - 15:00","status":"available"},{"key":"sat-15:00","label":"15:00 - 16:00","status":"available"},{"key":"sat-16:00","label":"16:00 - 17:00","status":"available"}]}]')
on conflict do nothing;

insert into trainer_clients (id, trainer_id, client_name, equipment_name, plan_name, next_session, contact, status)
values
  ('client-pim-1', 'coach-pim', 'Mina S.', 'Pilates Reformer', 'Progress 3 months', 'Monday 10:00', 'LINE: mina.homefit', 'active'),
  ('client-pim-2', 'coach-pim', 'Ploy T.', 'Stability Chair', 'Starter 1 month', 'Wednesday 14:00', 'LINE: ploycore', 'active'),
  ('client-tone-1', 'coach-tone', 'Ken A.', 'Functional Trainer', 'Signature 6 months', 'Tuesday 17:00', 'LINE: ken.fitlab', 'active')
on conflict do nothing;

insert into trainer_applications (id, name, email, phone, password_hash, specialty, machine_focus, status)
values
  ('trainer-application-1', 'Coach Praew', 'praew.trainer@reformrental.com', '0891112233', crypt('password123', gen_salt('bf')), 'Reformer rehab and posture reset', '{"Pilates Reformer","Cadillac / Tower"}', 'pending')
on conflict do nothing;

insert into rental_plans (id, name, months, discount, optional_sessions, required_sessions, note)
values
  ('starter', 'Starter 1 month', 1, 1.00, 2, 4, 'Best for short trials and space validation'),
  ('progress', 'Progress 3 months', 3, 0.94, 4, 12, 'The default balanced plan for long enough improvement'),
  ('signature', 'Signature 6 months', 6, 0.88, 8, 24, 'Long-term home studio style commitment')
on conflict do nothing;

insert into trainer_service_plans (id, name, sessions, discount, note)
values
  ('trainer-lite', 'Lite 4 Sessions', 4, 1.00, 'For light form checks when the client already has equipment'),
  ('trainer-core', 'Core 8 Sessions', 8, 0.95, 'The most likely trainer-only package for monthly coaching'),
  ('trainer-pro', 'Pro 12 Sessions', 12, 0.90, 'For clients who want close trainer follow-up')
on conflict do nothing;

insert into home_page_contents (id, payload)
values (
  'home-content-default',
  '{
    "hero": {
      "eyebrow": "Page 1 / Overview",
      "titleLine1": "Rent fitness equipment from home",
      "titleLine2": "Choose both equipment and trainer in one place",
      "description": "The first page explains the business clearly before the booking page handles package selection and pricing.",
      "primaryButtonLabel": "Go to booking page",
      "secondaryButtonLabel": "Rent reformer only"
    },
    "heroAside": {
      "badge": "Website Direction",
      "title": "Page 1 explains the business, page 2 handles real booking choices",
      "description": "This keeps the sales story clear before asking the client to choose a package."
    },
    "stats": [
      {
        "label": "Trainers available for hire",
        "description": "Coverage across reformer, rehab, and strength."
      },
      {
        "label": "Equipment available for rent",
        "description": "A mix of home-friendly and trainer-required machines."
      },
      {
        "value": "2 pages",
        "label": "Main website structure",
        "description": "Overview first, booking second."
      }
    ],
    "bookingModes": [
      {
        "id": "bundle",
        "title": "Rent equipment + trainer",
        "subtitle": "Best for clients who want the site to organize the full flow",
        "description": "Choose equipment, rental duration, and trainer in one flow."
      },
      {
        "id": "equipment-only",
        "title": "Rent equipment only",
        "subtitle": "For clients who want only the machine",
        "description": "Only home-friendly equipment can be rented without a trainer."
      },
      {
        "id": "trainer-only",
        "title": "Hire trainer only",
        "subtitle": "For clients who already own equipment",
        "description": "Choose trainer and session package without renting equipment."
      }
    ],
    "topRentals": {
      "eyebrow": "Top 3 Rentals",
      "title": "Most rented machines",
      "description": "Popular machines guide users into the booking page with preselected context.",
      "items": [
        {
          "equipmentId": "reformer",
          "rank": "01",
          "highlight": "Top rental",
          "summary": "Pilates Reformer is the easiest machine for home adoption."
        },
        {
          "equipmentId": "functional",
          "rank": "02",
          "highlight": "Strength favorite",
          "summary": "Functional Trainer is popular for guided strength programs."
        },
        {
          "equipmentId": "tower",
          "rank": "03",
          "highlight": "Popular in private rehab",
          "summary": "Cadillac / Tower is chosen often for corrective movement and rehab."
        }
      ]
    },
    "highlights": [
      {
        "title": "Home installation",
        "description": "The team measures space, delivers, installs, and prepares usage points."
      },
      {
        "title": "Trainer choice",
        "description": "Clients can choose the coach that matches their goals."
      },
      {
        "title": "Flexible package structure",
        "description": "Clients can choose full bundle, equipment only, or trainer only."
      }
    ],
    "process": {
      "eyebrow": "User Flow",
      "title": "Two-page website flow",
      "steps": [
        {
          "step": "01",
          "title": "Understand the business from page 1",
          "description": "The client learns the service model before booking."
        },
        {
          "step": "02",
          "title": "Choose the right flow on page 2",
          "description": "Bundle, equipment-only, or trainer-only."
        },
        {
          "step": "03",
          "title": "Review package and contact sales",
          "description": "The system calculates a quote before the sales team closes the deal."
        }
      ]
    },
    "cta": {
      "badge": "Ready to move",
      "title": "The next page turns this into a real package quote",
      "description": "Clients understand the offer before entering the detailed booking page.",
      "bundleLabel": "Start with full package",
      "bundleTitle": "Equipment + trainer",
      "bundleDescription": "Best for clients who want the site to organize everything.",
      "trainerLabel": "Start flexible",
      "trainerTitle": "Trainer only",
      "trainerDescription": "Best for clients who already own equipment."
    }
  }'::jsonb
)
on conflict do nothing;
