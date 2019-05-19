#! /usr/bin/env python3

import subprocess

# Choose Encryption or Decryption

answer = input("Encrypt[1] or Decrypt[2]?\n ")

# Encryption

if (answer == "1"):

    data = input("DATA FILE NAME:\n ")

    new_file = input("WHAT WOULD YOU LIKE TO NAME THE NEWLY ENCRYPTED FILE?\n ")

    pubkey = input("RECIPIENTS PUBLIC KEY:\n ")


    subprocess.run(["openssl", "rand","-out", "randompass.txt", "-base64", "120"])

    subprocess.run(["openssl", "enc", "-aes-256-cbc", "-salt", "-in", data, "-out",   new_file, "-pass", "file:randompass.txt"])

    subprocess.run(["openssl", "rsautl","-encrypt", "-in", "randompass.txt", "-out"  , "randompass.enc", "-pubin", "-inkey",  pubkey])

    
    subprocess.run(["rm", "randompass.txt"])
    subprocess.run(["mv", data, "msg"])

# Decryption

elif (answer == "2"):

   privkey = input("ENTER YOUR PRIVATE KEY:\n ")

   file_name = input("ENTER THE NAME OF THE ENCRYPTED FILE:\n ")

   data = input("NAME THE NEW FILE:\n ")

   subprocess.run(["openssl", "rsautl", "-decrypt", "-inkey", privkey, "-in", "randompass.enc","-out", "decodedpass.txt"])


   subprocess.run(["openssl","enc", "-d", "-aes-256-cbc", "-in", file_name,"-out",  data, "-pass", "file:decodedpass.txt"])
  
   subprocess.run(["rm", "decodedpass.txt",
"keys"])
   subprocess.run(["rm", "randompass.enc"])

else:
  print("Try Again. Input is case sensitive")
