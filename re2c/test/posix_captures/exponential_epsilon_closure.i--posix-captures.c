/* Generated by re2c */
// test for epsilon closure construction:
// exponential blowup if paths are not merged
// as soon as they arrive at the same NFA state

{
	YYCTYPE yych;
	yyt1 = YYCURSOR;
	{
		const size_t yynmatch = 3;
		const YYCTYPE *yypmatch[yynmatch * 2];
		yypmatch[0] = yyt1;
		yypmatch[3] = yyt1;
		yypmatch[5] = yyt1;
		yypmatch[1] = YYCURSOR;
		yypmatch[2] = yyt1;
		yypmatch[4] = yyt1;
		{}
	}
}

re2c: warning: line 5: rule matches empty string [-Wmatch-empty-string]
