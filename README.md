as-startup/apps - Something Something AI
----------------------------------------

This is an application monorepo for the AI Startup system.
AI Startup is the latest thing in AI innovation.
LLMs, GPTs, convolutional whatsa-ma-jigs.
You name it, we're probably selling it.

This this repository is organized by service.
Each directory under the `services` directory is a service built and shipped.

We practice continuous trunk-based deployments and GitOps methodologies.
Everything in `main` is destined for production.
The `ai-startup/config` repo is where we keep all of production deployment and release configuration.

In order to decouple deployments from releases AI Startup uses Flipt.
