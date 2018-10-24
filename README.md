# Serverless Golang Battlesnake

This starter snakes assumed a few things:

1. A bit of familiarity with Golang.

2. We'll be using the [serverless](https://serverless.com/) framework for serverless deployment to AWS.


## What you get out of the box

* `/` landing page (GET)
* `/start` (POST)
* `/move` (POST)
* `/end` (POST)
* `/ping` landing page (POST)

---

# Installation

## 1. Install Golang

If you haven't, try setting up the [Golang Starter Snake](https://github.com/battlesnakeio/starter-snake-go) that's official maintained by Battlesnake. There are instructions there that should get you to the point of running a basic Golang project.

## 2. Install [Serverless](https://serverless.com/)

 [Install the Serverless framework](https://serverless.com/blog/anatomy-of-a-serverless-app/#setup)

 * Start by getting the "hello world" example working from [this tutorial](https://serverless.com/blog/framework-example-golang-lambda-support/).

 * Note: Requires Python (or Python 3) to install the AWS CLI

 * Note: Requires setting up an AWS account and setting up an IAM user

---

# Usage

Assuming that you've installed the serverless framework and have connected it to an AWS account:

## Gather dependencies.

```
Make get
```

## Build/compile

```
Make build
```

## Deployment

```
sls deploy
```

# Resources/Examples
- [Tutorial: Serverless Framework example for Golang and Lambda](https://serverless.com/blog/framework-example-golang-lambda-support/)
- [Repo: "AWS Golang simple http endpoint" Example](https://github.com/serverless/examples/tree/master/aws-golang-simple-http-endpoint)
- [Go start snake (used for API)](https://github.com/battlesnakeio/starter-snake-go)
