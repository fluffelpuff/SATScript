# SATScript - Proof of Concept

## Die Idee:

<b>Die Idee besteht darin:</b> Ethereum ähnliche Smart Contracts in einem Layer-2-Netzwerk auf Bitcoin auszuführen. Hierbei bietet Lightning eine perfekte Grundlage für Transaktionen in der zweiten Ebene.</br>
Wofür könnte diese Lösung verwendet werden?

* Komplexe MulitSig Wallet Lösungen (Exchange, Bridges, etc...)
* Hoch dezentralisierte Tauschbörsen (Uniswap Äquivalent)
* Bitcoin gedeckte Sidechains, z.b. mittels eines Proof of Stakes (hierbei würde die Sidechain zum dritten Layer von Bitcoin werden)
* Weitere Möglichkeiten:
* * Wetten
* * Oracle System
* * Echtzeit Spiele

## Die verwendeten Protokolle / Techniken

* Bitcoin (Basis Protokoll / L1)
* Nostr (Backend / Container Kommunikation)

## Eine Eigene Programmiersprache (SATScript):

Damit es eine perfekt auf Bitcoin angepasste Lösung gibt, habe ich mich entschieden, eine eigene Programmiersprache namens SATScript zu entwickeln. Theoretisch habe ich mich dabei an Ethereum orientiert, jedoch erinnert die virtuelle Maschine für diese Sprache in der Praxis eher an die JVM. (Noch befindet sich dich VM in Arbeit, Status: Lexer)

### Beispiel:

```
pragma satscript ^0.0.1

export func test() (Bool) {
	return true
}
```

Dieser Code stellt eine Funktion mit dem Namen Test bereit, welche ein True oder False zurückgibt.


# Hinweis

Die Hauptarbeiten finden derzeit in ''core/scriptparse'' statt.
