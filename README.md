# Sync Scribe

Sync Scribe is a real-time, collaborative note-taking application that enables users to create, view, and edit notes simultaneously. With Sync Scribe, you can experience seamless and synchronous note-taking across multiple devices, making it perfect for team collaboration and personal productivity.

## Features

- **Real-time Collaboration**: Work on notes together with your team in real-time, seeing changes instantly as they happen.
- **Intuitive User Interface**: Enjoy a clean and user-friendly interface that makes note-taking a breeze.
- **Tag Management**: Organize your notes effectively by assigning tags for easy categorization and retrieval.
- **Personalized Note Collection**: Keep your notes private and accessible only to you, ensuring confidentiality and personal organization.
- **Extensibility**: Sync Scribe is built with extensibility in mind, allowing developers to add new features and customize the application to suit specific needs.

## Technologies Used

Sync Scribe is built using the following technologies:

- Frontend: React
- Backend: Go
- Real-time Communication: WebSocket, OT algorithm
- Database: MongoDB
- Deployment: AWS Elastic Beanstalk

## Getting Started

To start using Sync Scribe, simply visit our website at [https://syncscribe.com](https://syncscribe.com) and create an account. Once you're logged in, you can begin creating, editing, and collaborating on notes instantly.

To launch the app locally move backend/.ebextensions/docker-compose.yml > SyncScribe/docker-compose.yml

Then run from root: docker-compose up --build 
                    mongod --dbpath ./data/db


You will need to setup and host a mongoDB server for local implementation

## Roadmap

I have an exciting roadmap planned for Sync Scribe, with several features and enhancements in the pipeline:

Stay tuned for updates and new releases as we continue to improve and expand Sync Scribe.

## License

Sync Scribe is released under the [MIT License](LICENSE).

## Acknowledgements

- [React](https://reactjs.org/)
- [Go](https://golang.org/)
- [AWS DynamoDB](https://aws.amazon.com/dynamodb/)
- [AWS Elastic Beanstalk](https://aws.amazon.com/elasticbeanstalk/)