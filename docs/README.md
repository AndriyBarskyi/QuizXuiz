# Spaced Repetition Test App

## Overview
The Spaced Repetition Test App is a customizable platform designed to improve long-term learning retention through various spaced repetition testing methods. Users can create accounts, track their progress, and engage in personalized study sessions that follow evidence-based learning practices.

## Features
### 1. User Management
- **User Accounts**: Users can create accounts, saving preferences, progress, and test results.
- **Progress Tracking**: Track learning history, review dates, and success rates.
- **User Roles**: Admin, instructor, and student roles with varying permissions.

### 2. Test Creation and Customization
- **Custom Test Types**: 
    - Flashcards (Basic SRS)
    - Multiple-Choice Questions
    - Fill-in-the-Blank
    - True/False
- **Spaced Repetition Algorithm Customization**: Customizable repetition intervals (fixed or performance-based).
- **Content Import/Export**: Import and export question sets (CSV, JSON, Anki Decks).

### 3. Learning Strategies and Best Practices
- **Active Recall Integration**: Requires users to actively recall information, enhancing memory.
- **Interleaved and Retrieval Practice**: Mix topics and question types within sessions and offer regular retrieval.
- **Elaboration and Dual Coding**: Support notes and visual aids to reinforce learning.

### 4. Customization and Flexibility
- **Customizable Interface**: Users can adjust color themes, font sizes, and layout.
- **Custom Learning Schedules**: Tailor study schedules with specific review dates or times.
- **Personalized Learning Paths**: Users can create learning tracks based on goals and skill gaps.
- **Difficulty Adjustment**: Option to adjust question difficulty and repetition frequency.

### 5. Feedback and Analysis
- **Performance Analytics**: Review heatmaps, progress over time, upcoming reviews, and mastery levels.
- **Adaptive Learning**: Adjust question difficulty or review intervals based on performance.
- **Memory Decay Alerts**: Notify users about unreviewed content.

### 6. Community and Sharing
- **Question Set Sharing**: Users can share or sell custom question sets.
- **Collaboration Tools**: Multiple users can edit a test set.
- **Discussion Board**: For users to discuss concepts and share tips.

### 7. Learning Context
- **Real-World Applications**: Spaced repetition reminders for real-life learning goals.
- **Gamification**: Streaks, badges, and leaderboards.
- **Smart Notifications**: Personalized reminders based on user habits.

### 8. Integrations and Extensions
- **API Support**: API for connecting with other tools (e.g., calendar or learning apps).
- **Browser Extensions and Mobile App**: For on-the-go access.
- **Third-Party Integrations**: Connect with tools like Google Calendar and Pomodoro timers.

### 9. Accessibility and Globalization
- **Multiple Languages**: Multilingual content support.
- **Accessibility Features**: Keyboard navigation, screen reader compatibility, adjustable contrast.

## Tech Stack
- **Backend Framework**: Go with Gin for efficient routing.
- **Database**: PostgreSQL managed by GORM ORM.
- **Config Management**: Viper for managing environment variables and configurations.
- **Authentication**: JSON Web Tokens (JWT) for secure session management.
- **Password Security**: bcrypt for password hashing.

## Project Structure
```plaintext
spaced-repetition-backend/
├── config/          # Configuration files and database setup
├── controllers/     # API endpoint handlers
├── middleware/      # Request validation and authentication
├── models/          # Database models
├── routes/          # Route setup and grouping
├── services/        # Business logic (auth, hashing, etc.)
├── utils/           # Utility functions (error handling, JWT helper)
└── main.go          # Entry point
```

## Getting Started

### Prerequisites
- **Go** v1.16+
- **PostgreSQL** database
- Git

### Installation
1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/spaced-repetition-backend.git
   cd spaced-repetition-backend
   ```

2. **Set up environment variables:**  
   Create a `.env` file with the following variables:
   ```plaintext
   DB_HOST=localhost
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=spaced_repetition
   DB_PORT=5432
   JWT_SECRET=your_jwt_secret
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the server:**
   ```bash
   go run main.go
   ```
   The server will be available at `http://localhost:8080`.

### API Documentation
Refer to [docs/API_REFERENCE.md](./docs/API_REFERENCE.md) for details on available endpoints, including example requests and responses.

### Roadmap
1. **Initial backend setup** with user authentication, JWT sessions, and core CRUD operations.
2. **API improvements** to support spaced repetition intervals, customization, and dynamic test creation.
3. **Frontend integration** with React.js for a seamless user experience.
4. **Extensions and mobile app** to allow users access from any device.

## License
[MIT](https://choosealicense.com/licenses/mit/)