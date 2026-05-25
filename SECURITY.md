# Security Policy

## Supported versions

This project is pre-1.0 and follows semantic versioning. Security fixes are
released against the latest published version. Please upgrade to the most
recent release before reporting an issue.

| Version | Supported          |
| ------- | ------------------ |
| 0.1.x   | :white_check_mark: |
| < 0.1   | :x:                |

## Reporting a vulnerability

Please report security issues **privately** using GitHub's private
vulnerability reporting — do not open a public issue or PR for a suspected
vulnerability.

1. Go to the repository's **Security** tab:
   <https://github.com/davidbudnick/shadcn-templ/security>
2. Click **Report a vulnerability** to open a private advisory.
3. Include the affected component, a description, reproduction steps, and the
   potential impact.

You'll receive an acknowledgement and, once the report is triaged, updates on
the fix and disclosure timeline through the advisory thread.

Because this is a server-side rendering library, the most relevant concerns are
HTML/attribute injection or escaping issues in component output. When reporting,
note whether attacker-controlled data reaches a component parameter.
