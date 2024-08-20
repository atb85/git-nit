# GIT - NIT 🪲

**code review tool**

## Why?

Proper use of Pull Requests and code review are necessary for [compliance](https://github.blog/enterprise-software/governance-and-compliance/demonstrating-end-to-end-traceability-with-pull-requests/) in many cases. And legal stuff aside, it is generally good practice to integrate code review with pull requests, it just makes sense.

Often, it can be easy to be lax and avoid proper use of pull requests. By not conducting a thorough review, bugs can be merged into main and be deployed. Due to the [exponential increase in costs](https://www.functionize.com/blog/the-cost-of-finding-bugs-later-in-the-sdlc) of catching bugs later on in production, I thought a basic tool which adds **nits** into a pull request and checks whether these **nits** made it through review would be useful. This could help a team to analyse how diligently they are conducting code review, and provide even more evidence to support compliance through evidence of thorough code review.

## How?

There are three stages to the git nit pipeline
- The production of **nits** by the committer
- The combing for **nits** by the reviewer 
- The checking for **nits** in reviewed code (almost certainly to be run as a GitHub Action)

### Production methods

- [pre-commit](https://pre-commit.com)
- Command line tool
- GitHub Action (fired on pull request)

### Combing methods

This is the challenging choice, we need something which doesn't become a tedious process for the reviewer while also being involved enough to encourage proper use

- [Commit suggestions](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/reviewing-changes-in-pull-requests/incorporating-feedback-in-your-pull-request)
- leaving a `nit` comment which can remove the nit with an [action triggered upon review](https://docs.github.com/en/actions/writing-workflows/choosing-when-your-workflow-runs/events-that-trigger-workflows#pull_request_review_comment)  

### Checking methods

- GitHub action fired on [approval](https://docs.github.com/en/actions/writing-workflows/choosing-when-your-workflow-runs/events-that-trigger-workflows#running-a-workflow-when-a-pull-request-is-approved)


## What does a *nit* look like
- Something which can't be predicted (otherwise a ctrl+f could allow a way-round)
- Something which does not break the code
- Currently using a single line comment based on file type
- can potentially be obvious? hash of combination of commit hashes?
- 
