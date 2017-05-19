Deltakere:

Erdvik, Magnus

Lie, Eva Kristine

Nguyen, Philip

Tellefsen, Erlend Frøysnes

Van Dijk, Richard

Younas, Osman



                                                    ICA02 - Gruppe 11
 


                                                         Oppgave 1
Hvor mange prosesser som kjører på din datamaskin?
174
 
Hvor mange prosesser som kjører på din virtuelle server i nettskyen?
115
      	
Kan man gi et nøyaktig antall? Begrunn.
      	
Hvor mange av prosessene som “kjører”?
 Prossesser er alltid «on and off» og dette skjer ved oppstart, ettersom det blir gjort endringer på datamaskinen.
      	
Hvis de ikke kjører, hvilke tilstander befinner de seg da?
Prossessene befinner seg i en «sleeping» stilstand.
  
  
  //Screenshot til oppgave 1; https://imgur.com/a/5wLb1

 
 
Teamarbeid: Oppsummer alle data i en tabell i deres team-besvarelsen. Sammenlign deres platformer og diskuter forkjeller:
https://uiano-my.sharepoint.com/personal/magnue16_uia_no/_layouts/15/WopiFrame.aspx?docid=1ccf29b7606264bedb182abe77b3d057a&authkey=Ab1cJOqbXO2pOsArjFsJSCc&expiration=2018-02-11T14%3a37%3a43.000Z&action=view
 
 
Hvilke komponenter (både fysiske og abstrakte) i deres datasystemer er involvert i oppstart, administrasjon og avslutning av prosesser? Definer komponentene du nevner.
 
Oppstart: Fysisk er power knappen som sender strøm gjennom ledninger til strømforsyningen og videre til hovedkortet. På hovedkortet blir signalene videresendt til BIOS som sitter fysisk på hovedkortet, som har software integrert inni seg. BIOS har hovedoppgaven for å starte opp pc-en ved at hardware kan kommunisere med hverandre og å starte opp operativsystemet.
 
BIOS setter opp «broer» mellom hardware, og gir CPU adgang til all hardware. Den søker også gjennom tilkoblede enheter på pc-en og gjenkjenner disse.
Tester blir utført for å sjekke om det finnes feil på fysisk komponentene.
BIOS inneholder primitive lese funksjoner “drivere” som gjør det mulig for å lokalisere sine oppgaver til å utføre de neste stegene.
 
Etter testene er blitt utført for å sjekke om det er noe feil på hardware og broene i mellom hardware er sammenkoblet, så gir BIOS fra seg kontrollen til enheten som er valgt som boot devicen. Som standard er dette en av de fysiske harddiskene i maskinen. BIOS henter de første 512 bytes fra harddisken som er valgt. Disse bytes er kalt master boot record.
 
En liten del av denne mengden som kalles “the bootstrap code” har som oppgave å finne neste del/fil av start opp prosessen.
Den leter etter partisjonene på harddisken(e) og finner hvor operativsystemet ligger.
De første 3 bytes av MBR gir fra seg en JMP instruksjon som forteller CPU om neste steg av boot prosessen til den “aktive” partisjonen, så CPU kan ta for seg det videre steget. Bootloaderen kalt Bootmgr for de nyere windows systemene finnes på partisjonen og CPU tar for seg de siste stegene i boot prosessen. Når bootloaderen blir kjørt, laster den inn konfigurasjonen fra operativsystemet. Hvis brukeren har flere operativsystemer installer, så får man valget om disse i dette steget. Når operativsystemet er valgt, så laster den inn kernelen og all kontroll er nå gitt videre til operativsystemet.
 
 
Kilde : https://neosmart.net/wiki/mbr-boot-process/
 



                                                         Oppgave 2
 
Først så begynner vi med å opprette en fil som heter “hello.go” på server.

 
2. Så skriver vi “nano hello.go” for å kunne kode inni selve bash filen, og limer inn følgende hello.go kode.

 
3. Så fortsetter vi med trykke “ctrl + x” ved å lagre endringene, og dermed skrev vi inn kommandoen “go run hello.go” og fikk printet ut resultatet fra koden”



4. Så kompilerer vi til kjørbar windows fil ved hjelp av følgende kode: “GOOS=windows GOARCH= 386 go build hello.go”. Slik blir det opprettet en “hello.exe” fil på serveren.


 
5. Filen overføres til lokal windows maskin ved koden “scp -i <*.pem fil> <server navn@ip-adresse>:<fildestinasjon på server> <ønsket filplassering på lokal maskin>”. For eksempel:
 
 // Screenshots til oppgave 2; http://imgur.com/a/WJcgx // se på bilde tekstene for rekkefølge


 
Nå er programmet overført fra server til lokal maskin.
 
                                                         Oppgave 3
                                                         
                                                         
                                                                                                                 
skriv sum_tests_ og tilsvarende test-funksjoner TestSum for følgende typer i golang: uint32, int32, int64, float64 (følg eksemplene i filene) - lag tester, som produserer “FAIL” og forklar


- Testene produserer feil fordi summen av tallene er større enn hva typen klarer. For eksempel uint32 kan ikke ha tall mindre enn 0 og større enn 4294967295. Så om vi summerer to tall under 0  ved bruk av uint32 summer funksjonen vil testen produsere en FAIL: 

- Testen viser at minus tallene “overflows” og kan dermed ikke brukes med uint32.
 
Foreslå og gjerne implementer en løsning som gjør at bruken av deres pakke er trygg for brukerne.

- For å lage en trygg pakke gjorde vi det slik at pakken kun bruker float64 til å summere tall og gi resultat. float64 er den typen som tar mest mulige input, derfor vil den kunne summere alle tall innenfor de andre typene, itillegg til sin egen grense, som er på Slik:

I kommandolinjen

// Screenshot til oppgaven; http://imgur.com/a/Ch9oC
 
 
                                                          Oppgave 4
                                                          
- Se mappe "Oppgave 4"
I følgende mappe er det en eksempel på en benchmark test, vi har ikke implementert en grafisk forestilling av testen, men testen kan kjøres ved: go test //testfil



                                                          Oppgave 5
- Se mappe for oppgave 5                                                         
 

