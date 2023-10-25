# Awesome-Walking-Skeleton
Various walking skeleton solution (e.g. PHP, Python, Golang).

![image](https://github.com/ping-yee/Awesome-Walking-Skeleton/assets/65348108/9c69bbf3-08a2-488a-9293-72f00bf5fb3b)

## Solution List

This is the current progress list for this project and **contribution is welcome!**, please reference [Contributing](https://github.com/ping-yee/Awesome-Walking-Skeleton/edit/main/README.md#contributing).

| Language | Framework | Database          | Completed           | Contributor         |
|----------|-----------|-------------------|---------------------|--------------------|
| Golang   | Gin       | MySQL             | :white_check_mark:  | [Ping-yee](https://github.com/ping-yee) |
| PHP      | Laravel   | MySQL             | :construction:      | [Ping-yee](https://github.com/ping-yee) |
| Python   | Flask     | MySQL             | :construction:      | [Ping-yee](https://github.com/ping-yee) |
| Python   | Django    | MySQL             | :x:                 | :x: |

## What Is Walking Skeleton?

> A "walking skeleton" is an implementation of the thinnest possible slice of real functionality that we can automatically build, deploy, and test end-to-end.

A "Walking Skeleton" is the minimal implementation of a real feature that we can automatically build, deploy, and end-to-end test.

You should choose an actual product feature, and the smaller it is, the better – it should be simple but genuine. In the process of completing this feature, we need to:

1. Establish a rough architecture and make the most basic technical decisions to create the walking skeleton.
2. Write an end-to-end test to automate the testing of this feature.
3. Set up an automatic CI/CD (Continuous Integration/Continuous Deployment) pipeline with the following stages:
4. Build
5. End-To-End Test
6. Deploy

## Why Do We Need a Walking Skeleton?

> To free me from the busy manual tasks so that I can focus more on creating and exploring.

> If automation isn't established on Day 1, creativity and discovery will be hindered by mundane chores.

1. **Instant Feedback**
   - Once the Walking Skeleton is established, every feature we develop can be quickly deployed through automation, allowing the team to receive feedback from users or stakeholders early in the development process.
   - This helps in defining the system's basic requirements, features, and user experience and enables rapid adjustments and improvements.
2. **Architecture Validation**
   - The Walking Skeleton can help the team validate the overall system architecture.
   - It provides an opportunity to test the stack and communication between different features and components to ensure their correct integration.
3. **Risk Mitigation**
   - It can be used to validate technology choices, test system reliability and performance, and proactively discover and address potential issues.
4. **Incremental Development**
   - The Walking Skeleton aligns with the principles of agile development, supporting incremental development.
   - Teams can start with the Walking Skeleton and gradually expand and add new features, enabling rapid delivery and sustainable development.
5. **Promotion of Collaboration and Communication**
   - The Walking Skeleton provides a common foundation, fostering collaboration and communication within and outside the development team.
   - By showcasing the core components and functionality of the system, team members can better understand the overall system structure and engage in discussions and decision-making based on this foundation.

## Contributing

We warmly welcome contributions in various programming languages and frameworks. This enables a wide range of developers to use our library for rapid project construction. 🚀

### Contribution Guidelines

When you decide to contribute, please make sure your solution is following directory structure and includes the necessary components outlined below. 🗂️

### Directory Structure

For instance, if your solution targets backend development, it should be located within the `Backend` directory. Inside the folder for your chosen programming language, create new subfolders with the names of the framework and the database you are using. 🛠️

```
Backend/
├── PHP/
│ └── Laravel+MySQL/
│ ├── Codeigniter4+MySQL
│ └── ...
├── Python/
│ └── flask+MySQL/
├── ...
└── ...
```

### Required Components 

Your contribution should meet the following criteria:

1. Implement minimal business logic.
2. Include Unit Tests.
3. Incorporate End-to-End Tests.
4. Set up a Github Actions CI environment that successfully pass all the tests.

### Contributors List

We extend our heartfelt thanks to all the people who have contributed to this project. Below is a list of contributors:

<a href="https://github.com/ping-yee/Awesome-Walking-Skeleton/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=ping-yee/Awesome-Walking-Skeleton" />
</a>

Made with [contrib.rocks](https://contrib.rocks).
