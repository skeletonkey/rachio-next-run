name: golangci-lint

on:
  pull_request:
    branches:
      - "*"

jobs:
  golangci:
    name: lint
    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: v1.52.2