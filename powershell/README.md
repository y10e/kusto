# powershell
## How to use

import-Module ./KustoCmd.psm1 -function *   
Get-Sample

例:
```
> import-Module ./KustoCmd.psm1 -function *
> Get-Sample

GAC    Version        Location
---    -------        --------
False  v4.0.30319     C:\Users\yuyon\source\kusto\powershell\microsoft.azure.kusto.tools.5.0.8\tools\Kusto.Data.dll
Executing query: 'StormEvents | limit 5' with connection string: 'Data Source=https://help.kusto.windows.net;Initial Catalog=Samples;AAD Federated Security=True'



StartTime           EndTime             EpisodeId EventId State          EventType         InjuriesDirect InjuriesIndirect DeathsDirect DeathsIndirect
---------           -------             --------- ------- -----          ---------         -------------- ---------------- ------------ --------------
2007/12/30 16:00:00 2007/12/30 16:05:00     11749   64588 GEORGIA        Thunderstorm Wind              0                0            0              0
2007/12/20 7:50:00  2007/12/20 7:53:00      12554   68796 MISSISSIPPI    Thunderstorm Wind              0                0            0              0
2007/09/29 8:11:00  2007/09/29 8:11:00      11091   61032 ATLANTIC SOUTH Waterspout                     0                0            0              0
2007/09/20 21:57:00 2007/09/20 22:05:00     11078   60913 FLORIDA        Tornado                        0                0            0              0
2007/09/18 20:00:00 2007/09/19 18:00:00     11074   60904 FLORIDA        Heavy Rain                     0                0            0              
```

## 備考
`%APPDATA%\kusto\userTokenCache.data` に認証情報のキャッシュが保持されるため削除すると再度認証を要求される。このキャッシュは kusto explorer とも共有している。