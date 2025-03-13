#!/usr/bin/env bash

output=$(go run ./cmd/seclang-linter/seclang-linter.go run ./test/testdata/owasp-crs/*)

expected="
validating file ./test/testdata/owasp-crs/REQUEST-901-INITIALIZATION.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-905-COMMON-EXCEPTIONS.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-911-METHOD-ENFORCEMENT.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-913-SCANNER-DETECTION.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-920-PROTOCOL-ENFORCEMENT.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-921-PROTOCOL-ATTACK.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-922-MULTIPART-ATTACK.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-930-APPLICATION-ATTACK-LFI.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-931-APPLICATION-ATTACK-RFI.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-932-APPLICATION-ATTACK-RCE.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-933-APPLICATION-ATTACK-PHP.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-934-APPLICATION-ATTACK-GENERIC.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-941-APPLICATION-ATTACK-XSS.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-942-APPLICATION-ATTACK-SQLI.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-944-APPLICATION-ATTACK-JAVA.conf.........success!
validating file ./test/testdata/owasp-crs/REQUEST-949-BLOCKING-EVALUATION.conf.........success!
validating file ./test/testdata/owasp-crs/RESPONSE-950-DATA-LEAKAGES.conf.........success!
validating file ./test/testdata/owasp-crs/RESPONSE-951-DATA-LEAKAGES-SQL.conf.........success!
validating file ./test/testdata/owasp-crs/RESPONSE-952-DATA-LEAKAGES-JAVA.conf.........success!
validating file ./test/testdata/owasp-crs/RESPONSE-953-DATA-LEAKAGES-PHP.conf.........success!
validating file ./test/testdata/owasp-crs/RESPONSE-954-DATA-LEAKAGES-IIS.conf.........success!
validating file ./test/testdata/owasp-crs/RESPONSE-955-WEB-SHELLS.conf.........success!
validating file ./test/testdata/owasp-crs/RESPONSE-959-BLOCKING-EVALUATION.conf.........success!
validating file ./test/testdata/owasp-crs/RESPONSE-980-CORRELATION.conf.........success!
"

diff <(echo ${output}) <(echo ${expected})