# BuildersCourses — Launch Analysis & Course Strategy

## What This Is

A **"Build Your Own X" learning platform** (BuildersMTY) where students learn by implementing real systems from scratch — no videos, no copy-paste. Students write code in a Monaco editor, run tests, and progressively unlock modules. Every passing submodule = a git commit. On completion, they get a certificate and can claim their repo.

**Tech stack:** Next.js + Supabase + SharkAuth (GitHub OAuth) + Docker runners on Hetzner VPS + Modal (for GPU/training workloads).

---

## Current Course Inventory

| Course | Language | Status |
|---|---|---|
| HTTP Server | Go | **Complete** (6 modules, 15 submodules, full resources) |
| HTTP Server | C | Placeholder (empty) |
| Memory Allocator | — | Placeholder (empty) |
| Claude Code | — | Placeholder (empty) |
| GPT (Karpathy-style) | Python | **In progress** (Modal containers for training) |

---

## Course Architecture Advantage

The workspace architecture supports authoring a course once and ramping up versions in multiple languages fast using Claude Code. This means:

- Author the reference solution + tests + resources in one language
- Use Claude Code to generate stubs, adapt tests, and translate resources to other languages
- Ship each course in 2-4 languages with minimal extra effort
- Target languages per course: **Go, Python, TypeScript, Rust, C** (pick 2-3 per course based on fit)

---

## Course Catalog Strategy

### Audience Segments

| Segment | What they want | What hooks them |
|---|---|---|
| **Newbies** | Confidence, fundamentals, "I built something real" | Quick wins, guided paths, language basics |
| **Juniors** | Depth beyond tutorials, portfolio projects | Name-brand projects (Redis, Git, Docker) |
| **Mids** | Fill knowledge gaps, systems understanding | "How does X actually work?" curiosity |
| **Seniors** | Stay sharp, fun challenges, niche deep-dives | Hard problems, AI/ML, esoteric systems |

---

## Fundamentals Track — "Learn the Language"

Shorter courses (2-6 hours) that teach language fundamentals through small build exercises. Not "Build Your Own X" — more like "Learn X by Building Small Things." These serve as **on-ramps** to the main courses and catch newbies who aren't ready to jump into a full system build.

Each fundamentals course follows the same structure across languages, making it easy to generate variants with Claude Code.

| Course | Languages | Est. Hours | Modules |
|---|---|---|---|
| **Fundamentals: Go** | Go | 4-6h | Variables & types, Control flow, Functions, Structs & interfaces, Error handling, Goroutines basics, File I/O, Building a CLI tool |
| **Fundamentals: Python** | Python | 3-5h | Variables & types, Control flow, Functions & closures, Classes, Error handling, Iterators & generators, File I/O, Building a CLI tool |
| **Fundamentals: Rust** | Rust | 6-8h | Ownership & borrowing, Structs & enums, Pattern matching, Error handling (Result/Option), Traits, Lifetimes basics, File I/O, Building a CLI tool |
| **Fundamentals: TypeScript** | TypeScript | 3-5h | Types & interfaces, Functions & generics, Async/await & promises, Error handling, Modules, Node.js basics, Building a CLI tool |
| **Fundamentals: C** | C | 5-7h | Pointers & memory, Arrays & strings, Structs, Dynamic allocation (malloc/free), File I/O, Preprocessing & compilation, Building a CLI tool |
| **Fundamentals: Java** | Java | 4-6h | Types & OOP, Interfaces & generics, Collections, Error handling, Streams, Concurrency basics, File I/O, Building a CLI tool |

**Key design principle:** Every fundamentals course ends with building a small CLI tool that uses everything learned. The student walks away with something real, not just exercises.

**Progression:** Fundamentals courses recommend which "Build Your Own X" courses to take next based on the language learned.

---

## Full Course Catalog

### Tier 1 — Launch (ship these first)

| # | Course | Difficulty | Languages | Target | Est. Hours | Notes |
|---|---|---|---|---|---|---|
| 1 | **HTTP Server** | Intermediate | Go, Python, TS | Mids, Seniors | 6-20h | Already complete in Go |
| 2 | **Shell** | Beginner | Python, Go | Newbies, Juniors | 4-10h | Low barrier, everyone uses a terminal, great first "Build Your Own X" |
| 3 | **Redis** | Intermediate | Go, Python, Rust | Mids, Seniors | 8-20h | High name recognition, teaches networking + data structures |
| 4 | **Git** | Beginner-Intermediate | Python, Go, C | Juniors, Mids | 6-15h | Demystifies the tool devs use daily |
| 5 | **Fundamentals: Go** | Beginner | Go | Newbies | 4-6h | On-ramp to HTTP Server, Redis, etc. |
| 6 | **Fundamentals: Python** | Beginner | Python | Newbies | 3-5h | On-ramp to GPT, Shell, etc. |

**6 courses (including 2 fundamentals) = credible launch catalog.**

### Tier 2 — Fast follows (first month post-launch)

| # | Course | Difficulty | Languages | Target | Notes |
|---|---|---|---|---|---|
| 7 | **GPT** | Advanced | Python | Mids, Seniors | Karpathy-style from scratch. Modal containers for training. Flagship advanced course. |
| 8 | **SQLite** | Intermediate-Advanced | C, Go, Rust | Mids, Seniors | Deep systems understanding, massive appeal |
| 9 | **Docker** | Intermediate | Go | Juniors, Mids | Hype topic, containers are magic until you build one |
| 10 | **Load Balancer** | Intermediate | Go, Rust | Mids, Seniors | Natural follow-up to HTTP Server |
| 11 | **Fundamentals: Rust** | Beginner | Rust | Newbies | On-ramp to systems courses |
| 12 | **Fundamentals: TypeScript** | Beginner | TypeScript | Newbies | On-ramp to frontend/fullstack courses |

### Tier 3 — Community builders (months 2-3)

| # | Course | Difficulty | Languages | Target | Notes |
|---|---|---|---|---|---|
| 13 | **Regex Engine** | Intermediate | Python, Go, Rust | Mids, Seniors | Beloved by the "from scratch" crowd |
| 14 | **Compiler** (tiny language) | Advanced | Go, C, Rust | Seniors | Prestige course, senior magnet |
| 15 | **Auth Server** (OAuth2) | Intermediate | Go, TS | Juniors, Mids | Practical, security-aware devs love this |
| 16 | **TCP Stack** | Advanced | C, Rust | Seniors | Raw networking, deep systems |
| 17 | **DNS Resolver** | Intermediate | Go, Python, Rust | Juniors, Mids | Small scope, high "aha" factor |
| 18 | **Claude Code** | Advanced | Python, Go | Mids, Seniors | Build an AI coding agent. Unique offering nobody else has. |
| 19 | **Fundamentals: C** | Beginner | C | Newbies | On-ramp to Memory Allocator, SQLite, TCP Stack |

### Tier 4 — Differentiators & growth (months 3-6)

| # | Course | Difficulty | Languages | Target | Notes |
|---|---|---|---|---|---|
| 20 | **Memory Allocator** | Advanced | C | Seniors | Already planned. Deep systems. |
| 21 | **React** (mini framework) | Intermediate | TypeScript | Juniors, Mids | Frontend devs want this |
| 22 | **eBPF Tracer** | Advanced | C, Go | Seniors | Niche but generates buzz |
| 23 | **Message Queue** | Intermediate | Go, Rust | Mids, Seniors | Kafka-lite, teaches distributed systems |
| 24 | **JSON Parser** | Beginner | Any language | Newbies, Juniors | Tiny scope, great for first-timers. Can be a "warm-up" course. |
| 25 | **Blockchain** (simple) | Intermediate | Go, Python | Mids | Still has draw, teaches crypto + consensus |
| 26 | **Web Framework** | Intermediate | Go, Python, TS | Mids | Build Express/Gin from scratch, natural after HTTP Server |
| 27 | **Rate Limiter** | Beginner-Intermediate | Go, Python | Juniors, Mids | Small scope, practical, teaches algorithms (token bucket, sliding window) |
| 28 | **Key-Value Store** | Intermediate | Rust, Go | Mids, Seniors | LSM trees, WAL, compaction. Mini RocksDB. |
| 29 | **Container Runtime** | Advanced | Go | Seniors | Goes deeper than the Docker course — namespaces, cgroups, overlayfs |
| 30 | **Proxy Server** | Intermediate | Go, Rust | Mids, Seniors | Reverse proxy, TLS termination, practical infra knowledge |
| 31 | **Testing Framework** | Beginner-Intermediate | Python, Go, TS | Juniors, Mids | Build pytest/Jest from scratch. Meta and educational. |
| 32 | **Package Manager** | Intermediate | Go, TS | Mids | Dependency resolution, lockfiles, registries. Surprisingly deep. |
| 33 | **Debugger** | Advanced | C, Rust | Seniors | ptrace, DWARF, breakpoints. Niche but incredible learning. |
| 34 | **Image Format Decoder** (PNG) | Intermediate | C, Rust, Go | Mids, Seniors | Binary parsing, compression (zlib), visual output. Satisfying. |
| 35 | **Cron Scheduler** | Beginner-Intermediate | Go, Python | Juniors, Mids | Parsing cron expressions, scheduling, process management. Small and fun. |
| 36 | **Link Shortener** | Beginner | Go, Python, TS | Newbies, Juniors | Full stack mini-project: hashing, storage, HTTP. Good early course. |
| 37 | **Chat Server** | Intermediate | Go, TS | Juniors, Mids | WebSockets, concurrency, real-time. Students can demo it. |
| 38 | **BitTorrent Client** | Advanced | Go, Python, Rust | Seniors | Networking, binary protocols, peer-to-peer. Cult classic "build your own." |
| 39 | **Spell Checker** | Beginner | Python, Go | Newbies, Juniors | Tries, edit distance, dictionary loading. Algorithmic but approachable. |
| 40 | **Static Site Generator** | Beginner-Intermediate | Go, Python, TS | Juniors, Mids | Markdown parsing, templating, file I/O. Practical and shippable. |

---

## GPT Course — Special Infrastructure

The **Build Your Own GPT** course (based on Karpathy's video) requires GPU compute for training. Infrastructure plan:

- **Runtime:** Modal containers (serverless GPU)
- **Flow:** Student writes training code locally in the editor -> platform ships code to Modal -> Modal runs training on GPU -> streams logs/metrics back to browser
- **Test types needed:** New test type `training` that validates model outputs after training (e.g., loss below threshold, generated text quality)
- **Cost consideration:** GPU time is expensive. Options: (1) give each student a compute budget, (2) use small models/datasets that train in minutes, (3) offer pre-trained checkpoints for later modules so students aren't blocked by training time

### Suggested GPT Course Modules

1. **Tokenizer** — Build a BPE tokenizer from scratch
2. **Bigram Model** — Simplest language model, predict next character
3. **Self-Attention** — Implement scaled dot-product attention
4. **Transformer Block** — Multi-head attention + feed-forward + layer norm
5. **GPT Architecture** — Stack transformer blocks, add embeddings + positional encoding
6. **Training Loop** — Loss function, backprop, optimizer, training on tiny Shakespeare
7. **Text Generation** — Sampling, temperature, top-k, generating text from trained model

---

## Learning Paths (Recommended Tracks)

### Path: "I'm brand new to programming"
1. Fundamentals: Python (or Go)
2. JSON Parser
3. Shell
4. Link Shortener

### Path: "Backend Developer"
1. Fundamentals: Go
2. HTTP Server
3. Redis
4. Load Balancer
5. Web Framework
6. Auth Server

### Path: "Systems Programmer"
1. Fundamentals: C or Rust
2. Memory Allocator
3. SQLite
4. Container Runtime
5. TCP Stack
6. eBPF Tracer

### Path: "AI/ML Engineer"
1. Fundamentals: Python
2. GPT (Karpathy-style)
3. Claude Code (AI agent)

### Path: "Fullstack Developer"
1. Fundamentals: TypeScript
2. HTTP Server (TS)
3. React (mini framework)
4. Chat Server
5. Static Site Generator

### Path: "DevTools & Infrastructure"
1. Fundamentals: Go
2. Git
3. Docker
4. Testing Framework
5. Package Manager
6. Cron Scheduler

---

## UX Improvements Needed

### Onboarding & Discovery
- Add skill-level tags and difficulty indicators on course cards
- Add a "Start Here" badge on recommended first courses
- Show curated learning paths (tracks) on the landing page
- Show estimated time prominently on cards, adjusted to student's chosen difficulty

### Progress & Motivation
- **Streaks:** Daily coding streaks with visual indicator
- **Module badges:** Award badges per module completion (6 small wins per course, not just 1 at the end)
- **XP system:** Points for completing submodules, bonus for streaks, leaderboard
- **Public profiles:** Show completed builds, badges, streaks, languages used
- **Certificate improvements:** Module-level mini-certificates, not just course-level

### Stuck States
- "I'm stuck" button with progressive hint reveal: (1) nudge, (2) specific hint, (3) partial diff
- Optional AI tutor integration (Claude) — sees student's code + test failures, gives guidance without giving the answer
- Post-completion solution sharing — after passing a module, see how others solved it

### Social & Community
- Discussion threads per submodule
- "Show your solution" after completion (opt-in)
- Course completion leaderboard (fastest time, fewest hints used)
- Discord/community integration

### Multi-language
- Language picker on each course page
- "Also available in: Go, Python, Rust" badges
- Leverage Claude Code for fast language variant generation

### Mobile & Responsive
- Course catalog and progress tracking should be mobile-friendly
- Editor stays desktop-only with a clear "open on desktop to code" message
- Push notifications for streak reminders (mobile web)

### Search & Navigation
- Filter courses by: language, difficulty, topic, estimated time
- Search bar in catalog
- Tag system (networking, data structures, CLI, AI/ML, etc.)

---

## Launch Readiness Checklist

| Area | Status | Priority | Action |
|---|---|---|---|
| HTTP Server (Go) | Done | — | Ship as-is |
| Shell course | Not started | P0 | Author in Python, generate Go variant |
| Redis course | Not started | P0 | Author in Go, generate Python variant |
| Git course | Not started | P0 | Author in Python |
| Fundamentals: Go | Not started | P0 | Author (shorter, 4-6h) |
| Fundamentals: Python | Not started | P0 | Author (shorter, 3-5h) |
| GPT course | In progress | P1 | Complete, Modal infra needed |
| Platform (Next.js app) | Specced, not built | P0 | Build it |
| Docker runner pool | Specced, not built | P0 | Implement |
| Modal integration | Not started | P1 | Needed for GPT course |
| Learning paths UI | Not designed | P1 | Design & implement |
| Gamification (streaks, XP) | Not designed | P2 | Design & implement post-launch |
| AI tutor integration | Not designed | P2 | Design & implement post-launch |
| Multi-language variants | Architecture ready | P1 | Use Claude Code to generate |
