package timezone

var timezones = map[string][]string{
	"GMT": []string{
		"Africa/Abidjan",
		"Africa/Accra",
		"Africa/Bamako",
		"Africa/Banjul",
		"Africa/Bissau",
		"Africa/Conakry",
		"Africa/Dakar",
		"Africa/Freetown",
		"Africa/Lome",
		"Africa/Monrovia",
		"Africa/Nouakchott",
		"Africa/Ouagadougou",
		"Africa/Sao_Tome",
		"Africa/Timbuktu",
		"America/Danmarkshavn",
		"Antarctica/Troll",
		"Atlantic/Reykjavik",
		"Atlantic/St_Helena",
		"Eire",
		"Etc/GMT",
		"Etc/GMT+0",
		"Etc/GMT-0",
		"Etc/GMT0",
		"Etc/Greenwich",
		"Europe/Belfast",
		"Europe/Dublin",
		"Europe/Guernsey",
		"Europe/Isle_of_Man",
		"Europe/Jersey",
		"Europe/London",
		"GB",
		"GB-Eire",
		"GMT",
		"GMT+0",
		"GMT-0",
		"GMT0",
		"Greenwich",
		"Iceland",
	},
	"GHST": []string{
		"Africa/Accra",
	},
	"EAT": []string{
		"Africa/Addis_Ababa",
		"Africa/Asmara",
		"Africa/Asmera",
		"Africa/Dar_es_Salaam",
		"Africa/Djibouti",
		"Africa/Juba",
		"Africa/Kampala",
		"Africa/Mogadishu",
		"Africa/Nairobi",
		"Indian/Antananarivo",
		"Indian/Comoro",
		"Indian/Mayotte",
	},
	"CET": []string{
		"Africa/Algiers",
		"Africa/Ceuta",
		"Africa/Tunis",
		"Arctic/Longyearbyen",
		"Atlantic/Jan_Mayen",
		"CET",
		"Europe/Amsterdam",
		"Europe/Andorra",
		"Europe/Belgrade",
		"Europe/Berlin",
		"Europe/Bratislava",
		"Europe/Brussels",
		"Europe/Budapest",
		"Europe/Busingen",
		"Europe/Copenhagen",
		"Europe/Gibraltar",
		"Europe/Ljubljana",
		"Europe/Luxembourg",
		"Europe/Madrid",
		"Europe/Malta",
		"Europe/Monaco",
		"Europe/Oslo",
		"Europe/Paris",
		"Europe/Podgorica",
		"Europe/Prague",
		"Europe/Rome",
		"Europe/San_Marino",
		"Europe/Sarajevo",
		"Europe/Skopje",
		"Europe/Stockholm",
		"Europe/Tirane",
		"Europe/Vaduz",
		"Europe/Vatican",
		"Europe/Vienna",
		"Europe/Warsaw",
		"Europe/Zagreb",
		"Europe/Zurich",
		"Poland",
	},
	"WAT": []string{
		"Africa/Bangui",
		"Africa/Brazzaville",
		"Africa/Douala",
		"Africa/Kinshasa",
		"Africa/Lagos",
		"Africa/Libreville",
		"Africa/Luanda",
		"Africa/Malabo",
		"Africa/Ndjamena",
		"Africa/Niamey",
		"Africa/Porto-Novo",
	},
	"CAT": []string{
		"Africa/Blantyre",
		"Africa/Bujumbura",
		"Africa/Gaborone",
		"Africa/Harare",
		"Africa/Khartoum",
		"Africa/Kigali",
		"Africa/Lubumbashi",
		"Africa/Lusaka",
		"Africa/Maputo",
		"Africa/Windhoek",
	},
	"EET": []string{
		"Africa/Cairo",
		"Africa/Tripoli",
		"Asia/Amman",
		"Asia/Beirut",
		"Asia/Damascus",
		"Asia/Famagusta",
		"Asia/Gaza",
		"Asia/Hebron",
		"Asia/Nicosia",
		"EET",
		"Egypt",
		"Europe/Athens",
		"Europe/Bucharest",
		"Europe/Chisinau",
		"Europe/Helsinki",
		"Europe/Kaliningrad",
		"Europe/Kiev",
		"Europe/Kyiv",
		"Europe/Mariehamn",
		"Europe/Nicosia",
		"Europe/Riga",
		"Europe/Sofia",
		"Europe/Tallinn",
		"Europe/Tiraspol",
		"Europe/Uzhgorod",
		"Europe/Uzhhorod",
		"Europe/Vilnius",
		"Europe/Zaporozhye",
		"Europe/Zaporizhzhia",
		"Libya",
	},
	"EEST": []string{
		"Africa/Cairo",
		"Asia/Amman",
		"Asia/Beirut",
		"Asia/Damascus",
		"Asia/Famagusta",
		"Asia/Gaza",
		"Asia/Hebron",
		"Asia/Nicosia",
		"EET",
		"Egypt",
		"Europe/Athens",
		"Europe/Bucharest",
		"Europe/Chisinau",
		"Europe/Helsinki",
		"Europe/Kaliningrad",
		"Europe/Kiev",
		"Europe/Kyiv",
		"Europe/Mariehamn",
		"Europe/Nicosia",
		"Europe/Riga",
		"Europe/Sofia",
		"Europe/Tallinn",
		"Europe/Tiraspol",
		"Europe/Uzhgorod",
		"Europe/Uzhhorod",
		"Europe/Vilnius",
		"Europe/Zaporozhye",
		"Europe/Zaporizhzhia",
	},
	"WET": []string{
		"Africa/Casablanca",
		"Africa/El_Aaiun",
		"Atlantic/Canary",
		"Atlantic/Faeroe",
		"Atlantic/Faroe",
		"Atlantic/Madeira",
		"Europe/Lisbon",
		"Portugal",
		"WET",
	},
	"WEST": []string{
		"Africa/Casablanca",
		"Africa/El_Aaiun",
		"Atlantic/Canary",
		"Atlantic/Faeroe",
		"Atlantic/Faroe",
		"Atlantic/Madeira",
		"Europe/Lisbon",
		"Portugal",
		"WET",
	},
	"CEST": []string{
		"Africa/Ceuta",
		"Africa/Tunis",
		"Antarctica/Troll",
		"Arctic/Longyearbyen",
		"Atlantic/Jan_Mayen",
		"CET",
		"Europe/Amsterdam",
		"Europe/Andorra",
		"Europe/Belgrade",
		"Europe/Berlin",
		"Europe/Bratislava",
		"Europe/Brussels",
		"Europe/Budapest",
		"Europe/Busingen",
		"Europe/Copenhagen",
		"Europe/Gibraltar",
		"Europe/Ljubljana",
		"Europe/Luxembourg",
		"Europe/Madrid",
		"Europe/Malta",
		"Europe/Monaco",
		"Europe/Oslo",
		"Europe/Paris",
		"Europe/Podgorica",
		"Europe/Prague",
		"Europe/Rome",
		"Europe/San_Marino",
		"Europe/Sarajevo",
		"Europe/Skopje",
		"Europe/Stockholm",
		"Europe/Tirane",
		"Europe/Vaduz",
		"Europe/Vatican",
		"Europe/Vienna",
		"Europe/Warsaw",
		"Europe/Zagreb",
		"Europe/Zurich",
		"Poland",
	},
	"SAST": []string{
		"Africa/Johannesburg",
		"Africa/Maseru",
		"Africa/Mbabane",
	},
	"CAST": []string{
		"Africa/Khartoum",
	},
	"HAT": []string{
		"America/Adak",
		"America/Atka",
		"HST",
		"Pacific/Honolulu",
		"Pacific/Johnston",
		"US/Aleutian",
		"US/Hawaii",
	},
	"HAST": []string{
		"America/Adak",
		"America/Atka",
		"HST",
		"Pacific/Honolulu",
		"Pacific/Johnston",
		"US/Aleutian",
		"US/Hawaii",
	},
	"HADT": []string{
		"America/Adak",
		"America/Atka",
		"Pacific/Honolulu",
		"Pacific/Johnston",
		"US/Aleutian",
		"US/Hawaii",
	},
	"AKT": []string{
		"America/Anchorage",
		"America/Juneau",
		"America/Metlakatla",
		"America/Nome",
		"America/Sitka",
		"America/Yakutat",
		"US/Alaska",
	},
	"AKST": []string{
		"America/Anchorage",
		"America/Juneau",
		"America/Metlakatla",
		"America/Nome",
		"America/Sitka",
		"America/Yakutat",
		"US/Alaska",
	},
	"AKDT": []string{
		"America/Anchorage",
		"America/Juneau",
		"America/Metlakatla",
		"America/Nome",
		"America/Sitka",
		"America/Yakutat",
		"US/Alaska",
	},
	"AT": []string{
		"America/Anguilla",
		"America/Antigua",
		"America/Aruba",
		"America/Barbados",
		"America/Blanc-Sablon",
		"America/Curacao",
		"America/Dominica",
		"America/Glace_Bay",
		"America/Goose_Bay",
		"America/Grenada",
		"America/Guadeloupe",
		"America/Halifax",
		"America/Kralendijk",
		"America/Lower_Princes",
		"America/Marigot",
		"America/Martinique",
		"America/Moncton",
		"America/Montserrat",
		"America/Port_of_Spain",
		"America/Puerto_Rico",
		"America/Santo_Domingo",
		"America/St_Barthelemy",
		"America/St_Kitts",
		"America/St_Lucia",
		"America/St_Thomas",
		"America/St_Vincent",
		"America/Thule",
		"America/Tortola",
		"America/Virgin",
		"Atlantic/Bermuda",
		"Canada/Atlantic",
	},
	"AST": []string{
		"America/Anguilla",
		"America/Antigua",
		"America/Aruba",
		"America/Barbados",
		"America/Blanc-Sablon",
		"America/Curacao",
		"America/Dominica",
		"America/Glace_Bay",
		"America/Goose_Bay",
		"America/Grenada",
		"America/Guadeloupe",
		"America/Halifax",
		"America/Kralendijk",
		"America/Lower_Princes",
		"America/Marigot",
		"America/Martinique",
		"America/Moncton",
		"America/Montserrat",
		"America/Port_of_Spain",
		"America/Puerto_Rico",
		"America/Santo_Domingo",
		"America/St_Barthelemy",
		"America/St_Kitts",
		"America/St_Lucia",
		"America/St_Thomas",
		"America/St_Vincent",
		"America/Thule",
		"America/Tortola",
		"America/Virgin",
		"Asia/Aden",
		"Asia/Baghdad",
		"Asia/Bahrain",
		"Asia/Kuwait",
		"Asia/Qatar",
		"Asia/Riyadh",
		"Atlantic/Bermuda",
		"Canada/Atlantic",
	},
	"BRT": []string{
		"America/Araguaina",
		"America/Bahia",
		"America/Belem",
		"America/Fortaleza",
		"America/Maceio",
		"America/Recife",
		"America/Santarem",
		"America/Sao_Paulo",
		"Brazil/East",
	},
	"BRST": []string{
		"America/Araguaina",
		"America/Bahia",
		"America/Belem",
		"America/Fortaleza",
		"America/Maceio",
		"America/Recife",
		"America/Sao_Paulo",
		"Brazil/East",
	},
	"ART": []string{
		"America/Argentina/Buenos_Aires",
		"America/Argentina/Catamarca",
		"America/Argentina/ComodRivadavia",
		"America/Argentina/Cordoba",
		"America/Argentina/Jujuy",
		"America/Argentina/La_Rioja",
		"America/Argentina/Mendoza",
		"America/Argentina/Rio_Gallegos",
		"America/Argentina/Salta",
		"America/Argentina/San_Juan",
		"America/Argentina/San_Luis",
		"America/Argentina/Tucuman",
		"America/Argentina/Ushuaia",
		"America/Buenos_Aires",
		"America/Catamarca",
		"America/Cordoba",
		"America/Jujuy",
		"America/Mendoza",
		"America/Rosario",
	},
	"ARST": []string{
		"America/Argentina/Buenos_Aires",
		"America/Argentina/Catamarca",
		"America/Argentina/ComodRivadavia",
		"America/Argentina/Cordoba",
		"America/Argentina/Jujuy",
		"America/Argentina/La_Rioja",
		"America/Argentina/Mendoza",
		"America/Argentina/Rio_Gallegos",
		"America/Argentina/Salta",
		"America/Argentina/San_Juan",
		"America/Argentina/Tucuman",
		"America/Argentina/Ushuaia",
		"America/Buenos_Aires",
		"America/Catamarca",
		"America/Cordoba",
		"America/Jujuy",
		"America/Mendoza",
		"America/Rosario",
	},
	"PYT": []string{
		"America/Asuncion",
	},
	"PYST": []string{
		"America/Asuncion",
	},
	"ET": []string{
		"America/Atikokan",
		"America/Cancun",
		"America/Cayman",
		"America/Coral_Harbour",
		"America/Detroit",
		"America/Fort_Wayne",
		"America/Grand_Turk",
		"America/Indiana/Indianapolis",
		"America/Indiana/Marengo",
		"America/Indiana/Petersburg",
		"America/Indiana/Vevay",
		"America/Indiana/Vincennes",
		"America/Indiana/Winamac",
		"America/Indianapolis",
		"America/Iqaluit",
		"America/Jamaica",
		"America/Kentucky/Louisville",
		"America/Kentucky/Monticello",
		"America/Louisville",
		"America/Montreal",
		"America/Nassau",
		"America/New_York",
		"America/Nipigon",
		"America/Panama",
		"America/Pangnirtung",
		"America/Port-au-Prince",
		"America/Thunder_Bay",
		"America/Toronto",
		"Canada/Eastern",
		"EST",
		"EST5EDT",
		"Jamaica",
		"US/East-Indiana",
		"US/Eastern",
		"US/Michigan",
	},
	"EST": []string{
		"America/Atikokan",
		"America/Cancun",
		"America/Cayman",
		"America/Coral_Harbour",
		"America/Detroit",
		"America/Fort_Wayne",
		"America/Grand_Turk",
		"America/Indiana/Indianapolis",
		"America/Indiana/Marengo",
		"America/Indiana/Petersburg",
		"America/Indiana/Vevay",
		"America/Indiana/Vincennes",
		"America/Indiana/Winamac",
		"America/Indianapolis",
		"America/Iqaluit",
		"America/Jamaica",
		"America/Kentucky/Louisville",
		"America/Kentucky/Monticello",
		"America/Louisville",
		"America/Montreal",
		"America/Nassau",
		"America/New_York",
		"America/Nipigon",
		"America/Panama",
		"America/Pangnirtung",
		"America/Port-au-Prince",
		"America/Thunder_Bay",
		"America/Toronto",
		"Canada/Eastern",
		"EST",
		"EST5EDT",
		"Jamaica",
		"US/East-Indiana",
		"US/Eastern",
		"US/Michigan",
	},
	"CT": []string{
		"America/Bahia_Banderas",
		"America/Belize",
		"America/Chicago",
		"America/Costa_Rica",
		"America/El_Salvador",
		"America/Guatemala",
		"America/Havana",
		"America/Indiana/Knox",
		"America/Indiana/Tell_City",
		"America/Knox_IN",
		"America/Managua",
		"America/Matamoros",
		"America/Menominee",
		"America/Merida",
		"America/Mexico_City",
		"America/Monterrey",
		"America/North_Dakota/Beulah",
		"America/North_Dakota/Center",
		"America/North_Dakota/New_Salem",
		"America/Rainy_River",
		"America/Rankin_Inlet",
		"America/Regina",
		"America/Resolute",
		"America/Swift_Current",
		"America/Tegucigalpa",
		"America/Winnipeg",
		"Asia/Chongqing",
		"Asia/Chungking",
		"Asia/Harbin",
		"Asia/Kashgar",
		"Asia/Macao",
		"Asia/Macau",
		"Asia/Shanghai",
		"Asia/Taipei",
		"CST6CDT",
		"Canada/Central",
		"Canada/Saskatchewan",
		"Cuba",
		"Mexico/General",
		"PRC",
		"ROC",
		"US/Central",
		"US/Indiana-Starke",
	},
	"CST": []string{
		"America/Bahia_Banderas",
		"America/Belize",
		"America/Chicago",
		"America/Costa_Rica",
		"America/El_Salvador",
		"America/Guatemala",
		"America/Havana",
		"America/Indiana/Knox",
		"America/Indiana/Tell_City",
		"America/Knox_IN",
		"America/Managua",
		"America/Matamoros",
		"America/Menominee",
		"America/Merida",
		"America/Mexico_City",
		"America/Monterrey",
		"America/North_Dakota/Beulah",
		"America/North_Dakota/Center",
		"America/North_Dakota/New_Salem",
		"America/Rainy_River",
		"America/Rankin_Inlet",
		"America/Regina",
		"America/Resolute",
		"America/Swift_Current",
		"America/Tegucigalpa",
		"America/Winnipeg",
		"Asia/Chongqing",
		"Asia/Chungking",
		"Asia/Harbin",
		"Asia/Kashgar",
		"Asia/Macao",
		"Asia/Macau",
		"Asia/Shanghai",
		"Asia/Taipei",
		"Asia/Urumqi",
		"CST6CDT",
		"Canada/Central",
		"Canada/Saskatchewan",
		"Cuba",
		"Mexico/General",
		"PRC",
		"ROC",
		"US/Central",
		"US/Indiana-Starke",
	},
	"CDT": []string{
		"America/Bahia_Banderas",
		"America/Belize",
		"America/Chicago",
		"America/Costa_Rica",
		"America/El_Salvador",
		"America/Guatemala",
		"America/Havana",
		"America/Indiana/Knox",
		"America/Indiana/Tell_City",
		"America/Knox_IN",
		"America/Managua",
		"America/Matamoros",
		"America/Menominee",
		"America/Merida",
		"America/Mexico_City",
		"America/Monterrey",
		"America/North_Dakota/Beulah",
		"America/North_Dakota/Center",
		"America/North_Dakota/New_Salem",
		"America/Rainy_River",
		"America/Rankin_Inlet",
		"America/Resolute",
		"America/Tegucigalpa",
		"America/Winnipeg",
		"Asia/Chongqing",
		"Asia/Chungking",
		"Asia/Macao",
		"Asia/Macau",
		"Asia/Shanghai",
		"Asia/Taipei",
		"CST6CDT",
		"Canada/Central",
		"Cuba",
		"Mexico/General",
		"PRC",
		"ROC",
		"US/Central",
		"US/Indiana-Starke",
	},
	"ADT": []string{
		"America/Barbados",
		"America/Blanc-Sablon",
		"America/Glace_Bay",
		"America/Goose_Bay",
		"America/Halifax",
		"America/Martinique",
		"America/Moncton",
		"America/Puerto_Rico",
		"America/Thule",
		"Asia/Baghdad",
		"Atlantic/Bermuda",
		"Canada/Atlantic",
	},
	"AMT": []string{
		"America/Boa_Vista",
		"America/Campo_Grande",
		"America/Cuiaba",
		"America/Manaus",
		"America/Porto_Velho",
		"Asia/Yerevan",
		"Brazil/West",
	},
	"AMST": []string{
		"America/Boa_Vista",
		"America/Campo_Grande",
		"America/Cuiaba",
		"America/Manaus",
		"America/Porto_Velho",
		"Asia/Yerevan",
		"Brazil/West",
	},
	"COT": []string{
		"America/Bogota",
	},
	"COST": []string{
		"America/Bogota",
	},
	"MT": []string{
		"America/Boise",
		"America/Cambridge_Bay",
		"America/Chihuahua",
		"America/Creston",
		"America/Dawson_Creek",
		"America/Denver",
		"America/Edmonton",
		"America/Fort_Nelson",
		"America/Hermosillo",
		"America/Inuvik",
		"America/Mazatlan",
		"America/Ojinaga",
		"America/Phoenix",
		"America/Shiprock",
		"America/Whitehorse",
		"America/Yellowknife",
		"Canada/Mountain",
		"Canada/Yukon",
		"MST",
		"MST7MDT",
		"Mexico/BajaSur",
		"Navajo",
		"US/Arizona",
		"US/Mountain",
	},
	"MST": []string{
		"America/Boise",
		"America/Cambridge_Bay",
		"America/Chihuahua",
		"America/Creston",
		"America/Dawson_Creek",
		"America/Denver",
		"America/Edmonton",
		"America/Fort_Nelson",
		"America/Hermosillo",
		"America/Inuvik",
		"America/Mazatlan",
		"America/Ojinaga",
		"America/Phoenix",
		"America/Shiprock",
		"America/Whitehorse",
		"America/Yellowknife",
		"Canada/Mountain",
		"Canada/Yukon",
		"MST",
		"MST7MDT",
		"Mexico/BajaSur",
		"Navajo",
		"US/Arizona",
		"US/Mountain",
	},
	"MDT": []string{
		"America/Boise",
		"America/Cambridge_Bay",
		"America/Chihuahua",
		"America/Denver",
		"America/Edmonton",
		"America/Hermosillo",
		"America/Inuvik",
		"America/Mazatlan",
		"America/Ojinaga",
		"America/Phoenix",
		"America/Shiprock",
		"America/Yellowknife",
		"Canada/Mountain",
		"MST7MDT",
		"Mexico/BajaSur",
		"Navajo",
		"US/Arizona",
		"US/Mountain",
	},
	"VET": []string{
		"America/Caracas",
	},
	"GFT": []string{
		"America/Cayenne",
	},
	"PT": []string{
		"America/Dawson",
		"America/Ensenada",
		"America/Los_Angeles",
		"America/Santa_Isabel",
		"America/Tijuana",
		"America/Vancouver",
		"Canada/Pacific",
		"Mexico/BajaNorte",
		"PST8PDT",
		"US/Pacific",
	},
	"PST": []string{
		"America/Dawson",
		"America/Ensenada",
		"America/Los_Angeles",
		"America/Santa_Isabel",
		"America/Tijuana",
		"America/Vancouver",
		"Canada/Pacific",
		"Mexico/BajaNorte",
		"PST8PDT",
		"Pacific/Pitcairn",
		"US/Pacific",
	},
	"EDT": []string{
		"America/Detroit",
		"America/Fort_Wayne",
		"America/Grand_Turk",
		"America/Guayaquil",
		"America/Indiana/Indianapolis",
		"America/Indiana/Marengo",
		"America/Indiana/Petersburg",
		"America/Indiana/Vevay",
		"America/Indiana/Vincennes",
		"America/Indiana/Winamac",
		"America/Indianapolis",
		"America/Iqaluit",
		"America/Jamaica",
		"America/Kentucky/Louisville",
		"America/Kentucky/Monticello",
		"America/Louisville",
		"America/Montreal",
		"America/Nassau",
		"America/New_York",
		"America/Nipigon",
		"America/Pangnirtung",
		"America/Port-au-Prince",
		"America/Thunder_Bay",
		"America/Toronto",
		"Canada/Eastern",
		"EST5EDT",
		"Jamaica",
		"US/East-Indiana",
		"US/Eastern",
		"US/Michigan",
	},
	"ACT": []string{
		"America/Eirunepe",
		"America/Porto_Acre",
		"America/Rio_Branco",
		"Australia/Adelaide",
		"Australia/Broken_Hill",
		"Australia/Darwin",
		"Australia/North",
		"Australia/South",
		"Australia/Yancowinna",
		"Brazil/Acre",
	},
	"ACST": []string{
		"America/Eirunepe",
		"America/Porto_Acre",
		"America/Rio_Branco",
		"Australia/Adelaide",
		"Australia/Broken_Hill",
		"Australia/Darwin",
		"Australia/North",
		"Australia/South",
		"Australia/Yancowinna",
		"Brazil/Acre",
	},
	"PDT": []string{
		"America/Ensenada",
		"America/Los_Angeles",
		"America/Santa_Isabel",
		"America/Tijuana",
		"America/Vancouver",
		"Canada/Pacific",
		"Mexico/BajaNorte",
		"PST8PDT",
		"US/Pacific",
	},
	"WGT": []string{
		"America/Godthab",
		"America/Nuuk",
	},
	"WGST": []string{
		"America/Godthab",
		"America/Nuuk",
	},
	"ECT": []string{
		"America/Guayaquil",
	},
	"GYT": []string{
		"America/Guyana",
	},
	"BOT": []string{
		"America/La_Paz",
	},
	"BST": []string{
		"America/La_Paz",
		"Europe/Belfast",
		"Europe/Guernsey",
		"Europe/Isle_of_Man",
		"Europe/Jersey",
		"Europe/London",
		"GB",
		"GB-Eire",
		"Pacific/Bougainville",
	},
	"PET": []string{
		"America/Lima",
	},
	"PEST": []string{
		"America/Lima",
	},
	"PMST": []string{
		"America/Miquelon",
	},
	"PMDT": []string{
		"America/Miquelon",
	},
	"UYT": []string{
		"America/Montevideo",
	},
	"UYST": []string{
		"America/Montevideo",
	},
	"FNT": []string{
		"America/Noronha",
		"Brazil/DeNoronha",
	},
	"FNST": []string{
		"America/Noronha",
		"Brazil/DeNoronha",
	},
	"SRT": []string{
		"America/Paramaribo",
	},
	"CLT": []string{
		"America/Punta_Arenas",
		"America/Santiago",
		"Chile/Continental",
	},
	"CLST": []string{
		"America/Santiago",
		"Chile/Continental",
	},
	"EHDT": []string{
		"America/Santo_Domingo",
	},
	"EGT": []string{
		"America/Scoresbysund",
	},
	"EGST": []string{
		"America/Scoresbysund",
	},
	"NT": []string{
		"America/St_Johns",
		"Canada/Newfoundland",
	},
	"NST": []string{
		"America/St_Johns",
		"Canada/Newfoundland",
	},
	"NDT": []string{
		"America/St_Johns",
		"Canada/Newfoundland",
	},
	"AWT": []string{
		"Antarctica/Casey",
		"Australia/Perth",
		"Australia/West",
	},
	"AWST": []string{
		"Antarctica/Casey",
		"Australia/Perth",
		"Australia/West",
	},
	"DAVT": []string{
		"Antarctica/Davis",
	},
	"DDUT": []string{
		"Antarctica/DumontDUrville",
	},
	"MIST": []string{
		"Antarctica/Macquarie",
	},
	"MAWT": []string{
		"Antarctica/Mawson",
	},
	"NZT": []string{
		"Antarctica/McMurdo",
		"Antarctica/South_Pole",
		"NZ",
		"Pacific/Auckland",
	},
	"NZST": []string{
		"Antarctica/McMurdo",
		"Antarctica/South_Pole",
		"NZ",
		"Pacific/Auckland",
	},
	"NZDT": []string{
		"Antarctica/McMurdo",
		"Antarctica/South_Pole",
		"NZ",
		"Pacific/Auckland",
	},
	"ROTT": []string{
		"Antarctica/Palmer",
		"Antarctica/Rothera",
	},
	"SYOT": []string{
		"Antarctica/Syowa",
	},
	"VOST": []string{
		"Antarctica/Vostok",
	},
	"ALMT": []string{
		"Asia/Almaty",
		"Asia/Qostanay",
	},
	"ALMST": []string{
		"Asia/Almaty",
	},
	"ANAT": []string{
		"Asia/Anadyr",
	},
	"AQTT": []string{
		"Asia/Aqtau",
		"Asia/Aqtobe",
		"Asia/Atyrau",
	},
	"AQTST": []string{
		"Asia/Aqtobe",
	},
	"TMT": []string{
		"Asia/Ashgabat",
		"Asia/Ashkhabad",
	},
	"AZT": []string{
		"Asia/Baku",
	},
	"AZST": []string{
		"Asia/Baku",
	},
	"ICT": []string{
		"Asia/Bangkok",
		"Asia/Ho_Chi_Minh",
		"Asia/Phnom_Penh",
		"Asia/Saigon",
		"Asia/Vientiane",
	},
	"KRAT": []string{
		"Asia/Barnaul",
		"Asia/Krasnoyarsk",
		"Asia/Novokuznetsk",
	},
	"KGT": []string{
		"Asia/Bishkek",
	},
	"BNT": []string{
		"Asia/Brunei",
	},
	"IST": []string{
		"Asia/Calcutta",
		"Asia/Colombo",
		"Asia/Jerusalem",
		"Asia/Kolkata",
		"Asia/Tel_Aviv",
		"Eire",
		"Europe/Dublin",
		"Israel",
	},
	"YAKT": []string{
		"Asia/Chita",
		"Asia/Khandyga",
		"Asia/Yakutsk",
	},
	"YAKST": []string{
		"Asia/Chita",
		"Asia/Khandyga",
		"Asia/Yakutsk",
	},
	"CHOT": []string{
		"Asia/Choibalsan",
	},
	"CHOST": []string{
		"Asia/Choibalsan",
	},
	"BDT": []string{
		"Asia/Dacca",
		"Asia/Dhaka",
	},
	"BDST": []string{
		"Asia/Dacca",
		"Asia/Dhaka",
	},
	"TLT": []string{
		"Asia/Dili",
	},
	"GST": []string{
		"Asia/Dubai",
		"Asia/Muscat",
		"Atlantic/South_Georgia",
	},
	"TJT": []string{
		"Asia/Dushanbe",
	},
	"TSD": []string{
		"Asia/Dushanbe",
	},
	"HKT": []string{
		"Asia/Hong_Kong",
		"Hongkong",
	},
	"HKST": []string{
		"Asia/Hong_Kong",
		"Hongkong",
	},
	"HOVT": []string{
		"Asia/Hovd",
	},
	"HOVST": []string{
		"Asia/Hovd",
	},
	"IRKT": []string{
		"Asia/Irkutsk",
	},
	"IRKST": []string{
		"Asia/Irkutsk",
	},
	"TRT": []string{
		"Asia/Istanbul",
		"Europe/Istanbul",
		"Turkey",
	},
	"WIB": []string{
		"Asia/Jakarta",
		"Asia/Pontianak",
	},
	"WIT": []string{
		"Asia/Jayapura",
	},
	"IDT": []string{
		"Asia/Jerusalem",
		"Asia/Tel_Aviv",
		"Israel",
	},
	"AFT": []string{
		"Asia/Kabul",
	},
	"PETT": []string{
		"Asia/Kamchatka",
	},
	"PKT": []string{
		"Asia/Karachi",
	},
	"PKST": []string{
		"Asia/Karachi",
	},
	"NPT": []string{
		"Asia/Kathmandu",
		"Asia/Katmandu",
	},
	"KRAST": []string{
		"Asia/Krasnoyarsk",
	},
	"MYT": []string{
		"Asia/Kuala_Lumpur",
		"Asia/Kuching",
	},
	"MLAST": []string{
		"Asia/Kuala_Lumpur",
	},
	"BORTST": []string{
		"Asia/Kuching",
	},
	"MAGT": []string{
		"Asia/Magadan",
	},
	"MAGST": []string{
		"Asia/Magadan",
		"Asia/Srednekolymsk",
	},
	"WITA": []string{
		"Asia/Makassar",
		"Asia/Ujung_Pandang",
	},
	"PHT": []string{
		"Asia/Manila",
	},
	"PHST": []string{
		"Asia/Manila",
	},
	"NOVT": []string{
		"Asia/Novosibirsk",
		"Asia/Tomsk",
	},
	"OMST": []string{
		"Asia/Omsk",
	},
	"OMSST": []string{
		"Asia/Omsk",
	},
	"ORAT": []string{
		"Asia/Oral",
	},
	"KT": []string{
		"Asia/Pyongyang",
		"Asia/Seoul",
		"ROK",
	},
	"KST": []string{
		"Asia/Pyongyang",
		"Asia/Seoul",
		"ROK",
	},
	"QYZT": []string{
		"Asia/Qyzylorda",
	},
	"QYZST": []string{
		"Asia/Qyzylorda",
	},
	"MMT": []string{
		"Asia/Rangoon",
		"Asia/Yangon",
	},
	"SAKT": []string{
		"Asia/Sakhalin",
	},
	"UZT": []string{
		"Asia/Samarkand",
		"Asia/Tashkent",
	},
	"UZST": []string{
		"Asia/Samarkand",
		"Asia/Tashkent",
	},
	"KDT": []string{
		"Asia/Seoul",
		"ROK",
	},
	"SGT": []string{
		"Asia/Singapore",
		"Singapore",
	},
	"MALST": []string{
		"Asia/Singapore",
		"Singapore",
	},
	"SRET": []string{
		"Asia/Srednekolymsk",
	},
	"GET": []string{
		"Asia/Tbilisi",
	},
	"IRST": []string{
		"Asia/Tehran",
		"Iran",
	},
	"IRDT": []string{
		"Asia/Tehran",
		"Iran",
	},
	"BTT": []string{
		"Asia/Thimbu",
		"Asia/Thimphu",
	},
	"JST": []string{
		"Asia/Tokyo",
		"Japan",
	},
	"JDT": []string{
		"Asia/Tokyo",
		"Japan",
	},
	"ULAT": []string{
		"Asia/Ulaanbaatar",
		"Asia/Ulan_Bator",
	},
	"ULAST": []string{
		"Asia/Ulaanbaatar",
		"Asia/Ulan_Bator",
	},
	"VLAT": []string{
		"Asia/Ust-Nera",
		"Asia/Vladivostok",
	},
	"VLAST": []string{
		"Asia/Ust-Nera",
		"Asia/Vladivostok",
	},
	"YEKT": []string{
		"Asia/Yekaterinburg",
	},
	"YEKST": []string{
		"Asia/Yekaterinburg",
	},
	"AZOT": []string{
		"Atlantic/Azores",
	},
	"AZOST": []string{
		"Atlantic/Azores",
	},
	"CVT": []string{
		"Atlantic/Cape_Verde",
	},
	"FKT": []string{
		"Atlantic/Stanley",
	},
	"AET": []string{
		"Australia/ACT",
		"Australia/Brisbane",
		"Australia/Canberra",
		"Australia/Currie",
		"Australia/Hobart",
		"Australia/Lindeman",
		"Australia/Melbourne",
		"Australia/NSW",
		"Australia/Queensland",
		"Australia/Sydney",
		"Australia/Tasmania",
		"Australia/Victoria",
	},
	"AEST": []string{
		"Australia/ACT",
		"Australia/Brisbane",
		"Australia/Canberra",
		"Australia/Currie",
		"Australia/Hobart",
		"Australia/Lindeman",
		"Australia/Melbourne",
		"Australia/NSW",
		"Australia/Queensland",
		"Australia/Sydney",
		"Australia/Tasmania",
		"Australia/Victoria",
	},
	"AEDT": []string{
		"Australia/ACT",
		"Australia/Brisbane",
		"Australia/Canberra",
		"Australia/Currie",
		"Australia/Hobart",
		"Australia/Lindeman",
		"Australia/Melbourne",
		"Australia/NSW",
		"Australia/Queensland",
		"Australia/Sydney",
		"Australia/Tasmania",
		"Australia/Victoria",
	},
	"ACDT": []string{
		"Australia/Adelaide",
		"Australia/Broken_Hill",
		"Australia/Darwin",
		"Australia/South",
		"Australia/Yancowinna",
	},
	"ACWT": []string{
		"Australia/Eucla",
	},
	"ACWST": []string{
		"Australia/Eucla",
	},
	"ACWDT": []string{
		"Australia/Eucla",
	},
	"LHT": []string{
		"Australia/LHI",
		"Australia/Lord_Howe",
	},
	"LHST": []string{
		"Australia/LHI",
		"Australia/Lord_Howe",
	},
	"LHDT": []string{
		"Australia/LHI",
		"Australia/Lord_Howe",
	},
	"AWDT": []string{
		"Australia/Perth",
		"Australia/West",
	},
	"EAST": []string{
		"Chile/EasterIsland",
		"Pacific/Easter",
	},
	"EASST": []string{
		"Chile/EasterIsland",
		"Pacific/Easter",
		"Pacific/Galapagos",
	},
	"GMT-1": []string{
		"Etc/GMT+1",
	},
	"GMT-10": []string{
		"Etc/GMT+10",
	},
	"GMT-11": []string{
		"Etc/GMT+11",
	},
	"GMT-12": []string{
		"Etc/GMT+12",
	},
	"GMT-2": []string{
		"Etc/GMT+2",
	},
	"GMT-3": []string{
		"Etc/GMT+3",
	},
	"GMT-4": []string{
		"Etc/GMT+4",
	},
	"GMT-5": []string{
		"Etc/GMT+5",
	},
	"GMT-6": []string{
		"Etc/GMT+6",
	},
	"GMT-7": []string{
		"Etc/GMT+7",
	},
	"GMT-8": []string{
		"Etc/GMT+8",
	},
	"GMT-9": []string{
		"Etc/GMT+9",
	},
	"GMT+1": []string{
		"Etc/GMT-1",
	},
	"GMT+10": []string{
		"Etc/GMT-10",
	},
	"GMT+11": []string{
		"Etc/GMT-11",
	},
	"GMT+12": []string{
		"Etc/GMT-12",
	},
	"GMT+13": []string{
		"Etc/GMT-13",
	},
	"GMT+14": []string{
		"Etc/GMT-14",
	},
	"GMT+2": []string{
		"Etc/GMT-2",
	},
	"GMT+3": []string{
		"Etc/GMT-3",
	},
	"GMT+4": []string{
		"Etc/GMT-4",
	},
	"GMT+5": []string{
		"Etc/GMT-5",
	},
	"GMT+6": []string{
		"Etc/GMT-6",
	},
	"GMT+7": []string{
		"Etc/GMT-7",
	},
	"GMT+8": []string{
		"Etc/GMT-8",
	},
	"GMT+9": []string{
		"Etc/GMT-9",
	},
	"UTC": []string{
		"Etc/UCT",
		"Etc/UTC",
		"Etc/Universal",
		"Etc/Zulu",
		"UCT",
		"UTC",
		"Universal",
		"Zulu",
	},
	"SAMT": []string{
		"Europe/Astrakhan",
		"Europe/Samara",
		"Europe/Ulyanovsk",
	},
	"MSK": []string{
		"Europe/Kirov",
		"Europe/Minsk",
		"Europe/Moscow",
		"Europe/Simferopol",
		"W-SU",
	},
	"MSD": []string{
		"Europe/Kirov",
		"Europe/Moscow",
		"W-SU",
	},
	"GMT+04:00": []string{
		"Europe/Saratov",
	},
	"VOLT": []string{
		"Europe/Volgograd",
	},
	"-00": []string{
		"Factory",
	},
	"IOT": []string{
		"Indian/Chagos",
	},
	"CXT": []string{
		"Indian/Christmas",
	},
	"CCT": []string{
		"Indian/Cocos",
	},
	"TFT": []string{
		"Indian/Kerguelen",
	},
	"SCT": []string{
		"Indian/Mahe",
	},
	"MVT": []string{
		"Indian/Maldives",
	},
	"MUT": []string{
		"Indian/Mauritius",
	},
	"MUST": []string{
		"Indian/Mauritius",
	},
	"RET": []string{
		"Indian/Reunion",
	},
	"IRT": []string{
		"Iran",
	},
	"MHT": []string{
		"Kwajalein",
		"Pacific/Kwajalein",
		"Pacific/Majuro",
	},
	"MET": []string{
		"MET",
	},
	"MEST": []string{
		"MET",
	},
	"CHAT": []string{
		"NZ-CHAT",
		"Pacific/Chatham",
	},
	"CHAST": []string{
		"NZ-CHAT",
		"Pacific/Chatham",
	},
	"CHADT": []string{
		"NZ-CHAT",
		"Pacific/Chatham",
	},
	"WST": []string{
		"Pacific/Apia",
	},
	"WSDT": []string{
		"Pacific/Apia",
	},
	"CHUT": []string{
		"Pacific/Chuuk",
		"Pacific/Truk",
		"Pacific/Yap",
	},
	"VUT": []string{
		"Pacific/Efate",
	},
	"VUST": []string{
		"Pacific/Efate",
	},
	"PHOT": []string{
		"Pacific/Enderbury",
	},
	"TKT": []string{
		"Pacific/Fakaofo",
	},
	"FJT": []string{
		"Pacific/Fiji",
	},
	"FJST": []string{
		"Pacific/Fiji",
	},
	"TVT": []string{
		"Pacific/Funafuti",
	},
	"GALT": []string{
		"Pacific/Galapagos",
	},
	"GAMT": []string{
		"Pacific/Gambier",
	},
	"SBT": []string{
		"Pacific/Guadalcanal",
	},
	"ChST": []string{
		"Pacific/Guam",
		"Pacific/Saipan",
	},
	"GDT": []string{
		"Pacific/Guam",
		"Pacific/Saipan",
	},
	"LINT": []string{
		"Pacific/Kiritimati",
	},
	"KOST": []string{
		"Pacific/Kosrae",
	},
	"MART": []string{
		"Pacific/Marquesas",
	},
	"SST": []string{
		"Pacific/Midway",
		"Pacific/Pago_Pago",
		"Pacific/Samoa",
		"US/Samoa",
	},
	"NRT": []string{
		"Pacific/Nauru",
	},
	"NUT": []string{
		"Pacific/Niue",
	},
	"NFT": []string{
		"Pacific/Norfolk",
	},
	"NFDT": []string{
		"Pacific/Norfolk",
	},
	"NCT": []string{
		"Pacific/Noumea",
	},
	"NCST": []string{
		"Pacific/Noumea",
	},
	"PWT": []string{
		"Pacific/Palau",
	},
	"PONT": []string{
		"Pacific/Pohnpei",
		"Pacific/Ponape",
	},
	"PGT": []string{
		"Pacific/Port_Moresby",
	},
	"CKT": []string{
		"Pacific/Rarotonga",
	},
	"CKHST": []string{
		"Pacific/Rarotonga",
	},
	"TAHT": []string{
		"Pacific/Tahiti",
	},
	"GILT": []string{
		"Pacific/Tarawa",
	},
	"TOT": []string{
		"Pacific/Tongatapu",
	},
	"TOST": []string{
		"Pacific/Tongatapu",
	},
	"WAKT": []string{
		"Pacific/Wake",
	},
	"WFT": []string{
		"Pacific/Wallis",
	},
}
