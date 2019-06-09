import discord

client = discord.Client()
token = open("auth/token.txt", "r").readline()
guild_id = open("auth/guild_id.txt","r").readline()


@client.event
async def on_ready():
    print('We have logged in as {0.user}'.format(client))

@client.event
async def on_message(message):
    if message.author == client.user:
        return

@client.event
async def on_message(message):
    nullbot_guild = client.get_guild(int(guild_id.strip()))

    if "nb.count" == message.content.lower():
        await message.channel.send(f"```{nullbot_guild.member_count}```")

        

client.run(token.strip())
