# MeCP

An MCP for me, or anyone, giving context about useful personal details for LLMs.

## Example

"Write me an email to my car workshop to ask for a service, giving my car details so they know which parts to order."

## Data

Data is provided in a YAML file, which in a cloud native kubernetes environment could be mounted using a secret. It is not intended for highly person details, but basic details which help with everyday life can be added.
