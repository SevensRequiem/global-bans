static CGlobalVars *gpGlobals;
static IVEngineServer *engine;

CON_COMMAND( list_players, "Prints the name of each connected player" )
{
	for (int i=1;i<gpGlobals->maxClients;i++) // EntIndex 0 is worldspawn, after which come the players
	{
		// n.b. this only works if the player is active; players still connecting won't show up
		IPlayerInfo *playerinfo = playerinfomanager->GetPlayerInfo(engine->PEntityOfEntIndex(i));
		if (playerinfo)
		{
			Msg(playerinfo->GetName());
			Msg("\n");
		}
	}
}