/* Generated by re2c */

{
	YYCTYPE yych;
	if (YYLIMIT <= YYCURSOR) YYFILL(1);
	yych = *YYCURSOR;
	yyt1 = YYCURSOR;
	if (yych >= 0x01) goto yy3;
yy2:
	yynmatch = 2;
	yypmatch[2] = yyt1;
	yypmatch[0] = yyt1;
	yypmatch[1] = YYCURSOR;
	yypmatch[3] = YYCURSOR;
	{}
yy3:
	++YYCURSOR;
	goto yy2;
}

re2c: warning: line 5: rule matches empty string [-Wmatch-empty-string]
re2c: warning: line 6: rule matches empty string [-Wmatch-empty-string]
re2c: warning: line 6: unreachable rule (shadowed by rule at line 5) [-Wunreachable-rules]