# devops-demo-app

## tl;dr
This repo is one of three related repos:
- [IaC repo](https://github.com/MSFT-MCSEMEA-Demo/DevOpsDemo) - deploying relevant services to Azure
- [argoCD apps](https://github.com/MSFT-MCSEMEA-Demo/argocd-apps-config) - deploying argoCD apps 
- [DemoApp](https://github.com/MSFT-MCSEMEA-Demo/devops-demo-app) - deploying demoapp (Go)
  
The goal is to demonstrate a complete CI/CD pipeline using GitHub Actions implmenting DevSecOps practices using GitHub Advanced Security tools: Code Scanning (CodeQL), dependency review (depandadbot) and secret scanning. 
Additional security scanning which being used in the demo are:
1. IaC check using [Checkov](https://www.checkov.io/)
2. Container image scanning the demo using [Aqua Trivy](https://github.com/aquasecurity/trivy)

## solution diagram
![alt text](diagram/diagram.png)

## pre-requisites
1. Azure subscription
2. Under Settings of current GitHub repo - create three GitHub environments in GitHub: dev, prod, whatif. for prod environment setup a need for manual approval.
3. Configure OpenID Connect in Azure. After completing the instructions, the following secrets should be created in your repo secrets: AZURE_CLIENT_ID, AZURE_SUBSCRIPTION_ID, AZURE_TENANT_ID - you will need to add three entries for each 'Environment' types: dev, prod and whatif
4. Setup dev and main branches and use branch protection in GitHub to require Pull Request when target branch is: main

## repo review
1. root: go application with tests
1. .github/workflows: github actions files
  - codeql.yaml : run codeql - code scanning
  - dependency-review.yaml : run dependency review
  - onPull.yaml - trigger on-pull event for running the tests
  - onPush.yaml - trigger the CI process
  - trivy.yaml - trigger the aqua trivy container scan
