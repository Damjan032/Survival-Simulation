# Survival Simulation
Predmetni projekat iz naprednih tehnika programiranja
## Opis aplikacije

Survival Simulation prestavlja aplikaciju koja simulira preživljavanje populacije jedinki. Da bi jedinka peživela mora doći do hrane i eventualno se reprodukovala. 

## Članovi populacije
Članovi same populacije su inicijalno podeljeni na dva tipa, na 'dobre' i 'loše'. Dobra jedinka ako se na istom izvoru resursa nađe sa drugom jedinkom, dele hranu. Nasuprot dobrim, loši ako se susretnu sa drugom jedinkom, kradu od druge, čime joj smanjuju šansu za daljim preživljavanjem. 

## Pravila preživljavanja

Preživljavanje se odvija u iteracijama, koje su podeljene na dan i noć. Jedinke preko dana tragaju za hranom, dok se noću potencijalno reprodukuju ili umiru, ako nemaju dovoljno hrane.
Prilikom reprodukovanja moguće je da dođe do mutacija (promena tipa u odnosu na roditelja). Takođe, postoje grabljivice koje mogu da usmrte jedinku dok je u potrzi za hranom, sama hrana može biti otrovna.

U zavisnosti od količine hrane zavisi da li će jedinka preživeti, ali ako nema dovoljno hrane za sigurnu noć (šansa da preživi = 100%) njena šansa za reprodukcijom je 0. A svaki višak hrane povećava šansu sa povećavanjem populacije.

Moguće je da će se pravila proširivati i menjati u toku samog razvoja aplikacije.

<img src="https://user-images.githubusercontent.com/34902687/113480706-05e79380-9496-11eb-87e6-7cac6b665289.png" width="600" height="350">

## Arhitektura sistema

Sistem bi se sastojao iz dva dela: logičkog dela koji je implementiran uz pomoc Go programskog jezika i GUI-a koji bi bio implementiran u Pharou.

### Logička jedinica (Go)

Logička jedinica bi se sastojala od jednog Go procesa, smatram da bi u ovakvom načinu implementacije nema potrebe sa mikoserviskom arhitekturom.

Glavni zadaci logičke jedinice u toku dana su raspoređivanje samih jedinki i ishrana, dok bi noću na osnovu pravila i šansi obrađivala podatke i donosila zaključke o samoj jedinki.

### Korisnicki interfejs (Pharo)

Korisnički interfejs bi se realizovao uz pomoć Roassal biblioteke koja kao osnovu koristi Pharo.

Sastojao bi se iz dela za unos parametara i dela za vizuelizaciju i kontrolu iteracija.

Vizuelizacija bi bila u 2D tehnologiji i prikazivala bi promenu populacije u realnom vremenu.

## Dalja unapređenja

Neka od potencijalnih porširenja mogu biti:
- Paralelizacija logičke jedinice zasnovana na modelu deljene memorije (gouroutine)
- Upotreba logova koji bi omogućili retrospektivnu analizu podataka
- Proširenje same složenosti pravila
- Povećavanje broja tipova jedinki




