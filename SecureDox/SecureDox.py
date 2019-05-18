#! /usr/bin/env python3

import subprocess

data = input("What is the path of the data file you would like to encrypt? ")

new_file = input("What would you like to call the new file? ")

pubkey = input("What is the path of the public key? ")


subprocess.run(["openssl", "rand","-out", "newPass-b64.txt", "-base64", "120"])

subprocess.run(["openssl", "enc", "-aes-256-cbc", "-salt", "-in", data, "-out", new_file, "-pass", "file:newPass-b64.txt"])


subprocess.run(["openssl", "rsautl","-encrypt", "-in", "newPass-b64.txt", "-out", "newPass-b64.enc", "-pubin", "-inkey", pubkey])

subprocess.run(["mkdir", "SecureDox"])

subprocess.run(["mv", "newPass-b64.enc", new_file, "SecureDox"])


