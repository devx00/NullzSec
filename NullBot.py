import discord

client = discord.Client()

@client.event
async def on_ready():
    print('We have logged in as {0.user}'.format(client))

@client.event
async def on_message(message):
    if message.author == client.user:
        return

    if message.content.startswith('$hello'):
        await message.channel.send('Hello, Zach!')

client.run('NTg2ODAyNjc5MzYwMTI2OTg4.XPtVgQ.m5jLi7vEm1oS23im9dxIQX_SD0I')
