# Self-Debunking AI

## About

Self-Debunking AI is an application designed to critically analyze ideas you submit. Unlike typical feedback systems, this AI doesn't just offer encouragement - it provides logical, constructive criticism to help refine your concepts before development.

## Purpose

The goal isn't to destroy your enthusiasm, but to strengthen your ideas by identifying potential weaknesses, logical fallacies, and implementation challenges. By confronting these issues early, you can develop more robust and well-thought-out concepts.

## Features

- **Critical Analysis**: Receive detailed critique of your ideas based on logical reasoning
- **Constructive Feedback**: Get actionable suggestions for improvement
- **Bias Detection**: Identify potential blind spots in your thinking
- **Implementation Challenges**: Discover potential obstacles before investing resources

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: React/Vue/Angular (frontend implementation as submodule)
- **Database**: PostgreSQL
- **Deployment**: Docker containerization

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Git

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/self-debunking-ai.git
   cd self-debunking-ai
   ```

2. Set up environment variables:
   ```
   cp backend/.env.example backend/.env
   # Edit the .env file with your configuration
   ```

3. Start the application:
   ```
   docker-compose up -d
   ```

4. Access the application at http://localhost:3000

## Usage

1. Submit your idea through the web interface
2. The AI will analyze your concept for logical inconsistencies, implementation challenges, and potential improvements
3. Review the critique and use it to refine your idea
4. Iterate until your concept is solid and well-defined

## Project Structure

```
├── backend/             # Go backend API
├── frontend/            # Frontend application
├── docker-compose.yml   # Docker compose configuration
├── Dockerfile.backend   # Backend Docker configuration
└── Dockerfile.frontend  # Frontend Docker configuration
```

## License

KOCAN GANTENG

## Contributors

