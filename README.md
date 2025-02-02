# MMS (Message Management System)

A centralized platform for managing, tracking, and automating message distribution across different systems and teams in a data warehouse environment.

## 🎯 Purpose

MMS solves the challenge of scattered messaging scripts and undocumented communication processes by providing a unified platform for message management between upstream and downstream systems, as well as top-level reporting.

## ✨ Key Features

### Message Management
- Centralized message creation and editing
- Template management system
- Rich text editor with markdown support
- File attachment capabilities
- Message scheduling and automation
- Version control for messages
- Draft and approval workflow

### Routing & Distribution
- Multi-channel message delivery (Email, Slack, Teams, SMS)
- Configurable routing rules
- Recipient group management
- Priority levels and escalation
- Conditional delivery logic
- Bulk message handling

### Monitoring & Analytics
- Real-time delivery tracking
- Message status dashboard
- Delivery confirmation system
- Error tracking and alerts
- Performance metrics
- Historical data analysis
- Custom reporting

### Administration
- Role-based access control
- User management
- System configuration
- Integration management
- Audit logging
- Security settings

### API & Integration
- RESTful API
- Webhook support
- Database connectors
- Platform integrations (Slack, Teams, etc.)
- Custom script integration
- Batch processing capabilities

## 🛠️ Tech Stack

- **Frontend**: React + TypeScript
- **Backend**: Go (Gin framework)
- **Database**: PostgreSQL
- **Message Queue**: RabbitMQ
- **Caching**: Redis
- **Authentication**: JWT
- **Documentation**: Swagger/OpenAPI

## 📁 Project Structure
```
mms/
├── client/                 # React frontend
│   ├── src/
│   ├── public/
│   └── tests/
├── server/                 # Go backend
│   ├── cmd/
│   │   └── api/           # Application entrypoint
│   ├── internal/
│   │   ├── handlers/      # HTTP handlers
│   │   ├── models/        # Data models
│   │   ├── repository/    # Database operations
│   │   ├── services/      # Business logic
│   │   └── middleware/    # HTTP middleware
│   ├── pkg/               # Shared packages
│   │   ├── config/
│   │   ├── logger/
│   │   └── utils/
│   └── tests/             # Integration tests
├── docs/                   # Documentation
└── scripts/               # Utility scripts
```

## 🚀 Getting Started

### Prerequisites
- Go >= 1.21
- Node.js >= 18
- PostgreSQL >= 14
- RabbitMQ
- Redis

### Installation

1. Clone the repository
```bash
git clone https://github.com/pesnik/mms.git
cd mms
```

2. Install dependencies
```bash
# Install frontend dependencies
cd client
npm install

# Install Go dependencies
cd ../server
go mod download
```

3. Configure environment
```bash
# Copy environment files
cp .env.example .env
# Add your configuration values
```

4. Initialize database
```bash
# Run migrations
go run cmd/migrate/main.go
```

5. Start development servers
```bash
# Start backend (from server directory)
go run cmd/api/main.go

# Start frontend (from client directory)
npm run dev
```

## 📝 API Documentation

API documentation is available at `/swagger/index.html` when running the server locally.

## 🔐 Security Features

- JWT-based authentication
- Role-based access control
- Input validation
- Rate limiting
- Request logging
- SSL/TLS encryption
- Data encryption at rest

## 🧪 Testing

```bash
# Run backend tests
cd server
go test ./...

# Run frontend tests
cd client
npm test
```

## 📈 Roadmap

- [ ] Message template system
- [ ] Advanced routing rules
- [ ] Bulk message processing
- [ ] Custom integration framework
- [ ] Advanced analytics dashboard
- [ ] Mobile application
- [ ] Message archiving system
- [ ] AI-powered message optimization

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## 🙏 Acknowledgments

- Data Warehouse team for project requirements
- Contributors and testers
- Open source community

---
