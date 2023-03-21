# Call Billing Example

[![run tests](https://github.com/zenthangplus/call-billing-example/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/zenthangplus/call-billing-example/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/zenthangplus/call-billing-example/branch/main/graph/badge.svg)](https://codecov.io/gh/zenthangplus/call-billing-example)
[![goreportcard](https://goreportcard.com/badge/github.com/zenthangplus/call-billing-example)](https://goreportcard.com/report/github.com/zenthangplus/call-billing-example)

## Architecture design

> Abbreviations:
> - `CallCtl` means `CallController`
> - `BillingCtl` means `BillingController`
> - `AggrListener` means `BillingAggregationListener`

     ┌──────┐              ┌───────┐                ┌──┐          ┌────────┐          ┌────────────┐          ┌──────────┐
     │Client│              │CallCtl│                │DB│          │EventBus│          │AggrListener│          │BillingCtl│
     └──┬───┘              └───┬───┘                └┬─┘          └───┬────┘          └─────┬──────┘          └────┬─────┘
        │ 1.1. End call request│                     │                │                     │                      │
        │ ─────────────────────>                     │                │                     │                      │
        │                      │                     │                │                     │                      │
        │                      │ 1.2. Store call data│                │                     │                      │
        │                      │  to bills table     │                │                     │                      │
        │                      │ ────────────────────>                │                     │                      │
        │                      │                     │                │                     │                      │
        │                      │            1.3. Publish              │                     │                      │
        │                      │            CallEnded event           │                     │                      │
        │                      │ ─────────────────────────────────────>                     │                      │
        │                      │                     │                │                     │                      │
        │     1.4. Response    │                     │                │                     │                      │
        │ <─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─                     │                │                     │                      │
        │                      │                     │                │                     │                      │
        │                      │                     │                │   2.1. Subscribe    │                      │
        │                      │                     │                │   CallEnded event   │                      │
        │                      │                     │                │ <─ ─ ─ ─ ─ ─ ─ ─ ─ ─                       │
        │                      │                     │                │                     │                      │
        │                      │                     │      2.2. Aggregate billing and      │                      │
        │                      │                     │      store to billings table         │                      │
        │                      │                     │ <─────────────────────────────────────                      │
        │                      │                     │                │                     │                      │
        │                      │                  3.1. Get billing request                  │                      │
        │ ─────────────────────────────────────────────────────────────────────────────────────────────────────────>
        │                      │                     │                │                     │                      │
        │                      │                     │             3.2. Get billing from billings table            │
        │                      │                     │ <────────────────────────────────────────────────────────────
        │                      │                     │                │                     │                      │
        │                      │               3.3. Response with billing info              │                      │
        │ <─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─
     ┌──┴───┐              ┌───┴───┐                ┌┴─┐          ┌───┴────┐          ┌─────┴──────┐          ┌────┴─────┐
     │Client│              │CallCtl│                │DB│          │EventBus│          │AggrListener│          │BillingCtl│
     └──────┘              └───────┘                └──┘          └────────┘          └────────────┘          └──────────┘

## Setup

### Using docker

Simply run the following command to start all necessary modules:

```shell
docker-compose up
```
