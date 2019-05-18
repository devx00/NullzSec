Purpose:
The purpose of this script is to aid in administering a hybrid crypro-system.

Encryption:
This script utilizes the openssl tool set to produce a random 120 byte passphrase that is then used to encrypt a large data file using AES 256 Bit encryption. The passphrase is then encrypted with the recipients public key using RSA encryption. This allows a secure method of transfer of  both the encrypted data file and encrypted passphrase to the recipient. The recipient may then decrypt the passphrase using his/her RSA private key, and then in turn decrypt the data file using the passphrase.

Decryption:
Upon recieving the encrypted data file and encrypted passphrase from the sender, the user can use this script to decode the passphrase with his/her RSA private key.The newly decrypted passphrase is then used to decrypt the data file allowing a secure method of data transfer.

Benefits:
Not only does this script save time and effort in the command terminal, it is also a very secure means of handling both the data and passphrase.
This is because a newly randomized 120 byte passphrase is created with each use. Not only is this passphrase very strong, but it also defends against one of the biggest vulnerabilities to the AES-256-CBC passphrase system which is a keylogger. This is because the passphrase is never entered manually by neither the sender or recipient.
