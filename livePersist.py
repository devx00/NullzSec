#! /usr/bin/env python3
import subprocess


subprocess.run(["fdisk", "-l"])
print("\n\n\n")

device = input("What is the name of the storage device? (Ex. sda, sdb, sdc): ")


# Create Storage partition
subprocess.run(["fdisk", "/dev/" + device])
# Encrypt persistent storage

subprocess.run(["fdisk", "-l"])

part = input("What is the name of the storage partition? (Ex. sdb1, sdb2): ")

subprocess.run(["cryptsetup", "-v", "-y", "luksFormat", "/dev/" + part])

subprocess.run(["cryptsetup", "luksOpen", "/dev/" + part, "my_usb"])

subprocess.run(["mkfs.ext4", "-L persistence", "/dev/mapper/my_usb"])

subprocess.run(["e2label", "/dev/mapper/my_usb", "persistence"])

subprocess.run(["mkdir", "-p", "/mnt/my_usb"])

subprocess.run(["mount", "/dev/mapper/my_usb", "/mnt/my_usb"])

subprocess.run(["echo", "\"/ union \" > /mnt/my_usb/persistence.conf'])
