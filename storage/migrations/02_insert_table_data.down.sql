BEGIN;

	DELETE FROM "book" WHERE id = '433d6ba4-3d1d-4256-80d7-830bd6b30cc1';
	DELETE FROM "book" WHERE id = '433d6ba4-3d1d-4256-80d7-830bd6b30cc2';
	DELETE FROM "book" WHERE id = '433d6ba4-3d1d-4256-80d7-830bd6b30cc3';
	DELETE FROM "book" WHERE id = '433d6ba4-3d1d-4256-80d7-830bd6b30cc4';
	DELETE FROM "book" WHERE id = '433d6ba4-3d1d-4256-80d7-830bd6b30cc5';
	DELETE FROM "book" WHERE id = '433d6ba4-3d1d-4256-80d7-830bd6b30cc6';
	DELETE FROM "book" WHERE id = '433d6ba4-3d1d-4256-80d7-830bd6b30cc7';
	DELETE FROM "book" WHERE id = '433d6ba4-3d1d-4256-80d7-830bd6b30cc8';
	
	DELETE FROM "author" WHERE id='b9a39905-d3d4-4025-8905-21aa625a1aef';
	DELETE FROM "author" WHERE id='b9a39905-d3d4-4025-8905-22aa625a1aef';
	DELETE FROM "author" WHERE id='b9a39905-d3d4-4025-8905-23aa625a1aef';
	DELETE FROM "author" WHERE id='b9a39905-d3d4-4025-8905-24aa625a1aef';
	DELETE FROM "author" WHERE id='b9a39905-d3d4-4025-8905-25aa625a1aef';
	DELETE FROM "author" WHERE id='b9a39905-d3d4-4025-8905-26aa625a1aef';
	DELETE FROM "author" WHERE id='b9a39905-d3d4-4025-8905-27aa625a1aef';

	DELETE FROM "category" WHERE id='4b19a81c-f4b5-482b-bcc4-7d517a5cfd50';
	DELETE FROM "category" WHERE id='4b19a81c-f4b5-482b-bcc4-7d527a5cfd50';
	DELETE FROM "category" WHERE id='4b19a81c-f4b5-482b-bcc4-7d537a5cfd50';
	DELETE FROM "category" WHERE id='4b19a81c-f4b5-482b-bcc4-7d547a5cfd50';
	DELETE FROM "category" WHERE id='4b19a81c-f4b5-482b-bcc4-7d557a5cfd50';
	DELETE FROM "category" WHERE id='4b19a81c-f4b5-482b-bcc4-7d567a5cfd50';
	DELETE FROM "category" WHERE id='4b19a81c-f4b5-482b-bcc4-7d577a5cfd50';

	DELETE FROM "location" WHERE id='1319540f-fe91-4f8b-8bbe-93cfae5ed381';
	DELETE FROM "location" WHERE id='1319540f-fe92-4f8b-8bbe-93cfae5ed381';
	DELETE FROM "location" WHERE id='1319540f-fe93-4f8b-8bbe-93cfae5ed381';
	DELETE FROM "location" WHERE id='1319540f-fe94-4f8b-8bbe-93cfae5ed381';
	DELETE FROM "location" WHERE id='1319540f-fe95-4f8b-8bbe-93cfae5ed381';
	DELETE FROM "location" WHERE id='1319540f-fe96-4f8b-8bbe-93cfae5ed381';
	DELETE FROM "location" WHERE id='1319540f-fe97-4f8b-8bbe-93cfae5ed381';

COMMIT;



