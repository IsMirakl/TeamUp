# TeamUp
TeamUp - platform for finding partners and teams for startups and pet projects. Service helps find people based on skills, interests, and goals.

---

## Features
- User registration and authentication
- Profiles with skill, stacks and descriptions
- Creating projects and finding participants
- Chat between participants
- Notifications

---

## Tech Stack

### Frontend
- React 19
- TypeScript 5.9
- TailwindCSS 4.1
- Vite 7.3
- Zustand
- Axios

### Backend
- NestJS 11
- TypeScript 5.9
- PostgreSQL 18
- Prisma ORM 7.2

### Infrastructure
- Docker 29.1
- Docker Compose 5.0

---

## Getting Started

### 1 Clone repository
```bash
git clone https://github.com/IsMirakl/TeamUp.git
cd teamup
```
### 2 Start database
```bash
docker-compose up -d
```
### 3 Start Backend
```bash
cd backend
npm install
npx prisma migrate dev
npm run start:dev
```
### Start Frontend
```bash
cd frontend
npm install
npm run dev
```

---

# Api Documentation
After launch, the backend is available at:
http://localhost:3000/api

# Database
PostgreSQL is used as a database, and work with the database is carried out through Prisma ORM.
The database schema is described in:
docs/database.md

---

# Author
Daniil Bodrov
