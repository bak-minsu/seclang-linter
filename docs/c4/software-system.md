# System Context diagram

```mermaid
C4Context
    Enterprise_Boundary(enterprise, "Enterprise using Coraza WAF") {
        Person(admin, "Admin", "Administrator who wishes to find syntax errors in Coraza's SecLang")

        System(ci, "CI Automation", "Continuous Integration automation that checks for syntax errors in Coraza's SecLang")

        System_Boundary(system, "SecLang-Linter") {
            System(linter, "seclang-linter", "A linter which finds syntax errors in Coraza's SecLang")
        }

        Rel(admin, linter, "Runs")
        Rel(ci, linter, "Runs")
    }

    UpdateElementStyle(enterprise, $fontColor="white", $borderColor="white")
    UpdateElementStyle(system, $fontColor="white", $borderColor="white")

    UpdateRelStyle(admin, linter, $textColor="white", $lineColor="red", $offset="20")
    UpdateRelStyle(ci, linter, $textColor="white", $lineColor="red", $offset="20")
```