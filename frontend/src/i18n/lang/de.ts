export default {
    header: {
        server: {
            Tranquility: "Tranquility",
            Serenity: "Serenity"
        },
        region: {
            TheForge: "The Forge",
            Domain: "Domain",
            Heimatar: "Heimatar",
            SinqLaison: "Sinq Laison"
        }
    },
    config: {
        form: {
            materialPrice: "Materialpreis",
            materialPlaceholder: "Bitte wählen Sie einen Typ aus",
            productPrice: "Produktpreis",
            productPlaceholder: "Bitte wählen Sie einen Typ aus",
            scope: "Gewichteter Preis",
            scopePlaceholder: "Bitte wählen Sie einen Prozentsatz aus",
            days: "Datumsbereich",
            daysPlaceholder: "Bitte wählen Sie einen Datumsbereich aus",
            tax: "Steuern",
            week: "7 Tage",
            month: "30 Tage",
            season: "90 Tage",
            buyPrice: "Kaufpreis",
            sellPrice: "Verkaufspreis",
        },
        tableForm:{
            title: "Tabelle",
            calculator: "Taschenrechner",
            calculatorPlaceholder: "Taschenrechner aktivieren oder nicht",
            calculatorEnable: "Aktivieren",
            calculatorDisable: "Deaktivieren",
            layout: "Tabellenlayout",
            layoutPlaceholder: "Bitte Tabellenlayout auswählen",
            layoutAuto: "Automatisch",
            layoutFixed: "Fest",
            highlight: "Zeile hervorheben",
            highlightPlaceholder: "Hervorhebung der ausgewählten Zeile aktivieren oder nicht",
            highlightEnable: "Aktivieren",
            highlightDisable: "Deaktivieren",
        },
        desc: {
            title: "Beschreibung",
            dataDesc: "Daten",
            tableDesc: "Spalte",
            claim: "Copyright Hinweis",
            blueprint: "Bauplan",
            materialContent: "Preistyp für alle Materialien (Anforderungen, Bauplanmaterialien)",
            productPriceContent: "Preistyp für alle LP-Shop-Artikel",
            scopeContent: "Preisberechnung für alle Artikel. Niedrigpreisige und hochvolumige Bestellungen werden bereits im Backend herausgefiltert. Die Kauf- und Verkaufspreise sind Durchschnittspreise der niedrigsten oder höchsten x% Mengenpreise",
            daysContent: "Die Volumenspalte in der Tabelle stellt den Durchschnitt des Gesamtvolumens während des Datumsbereichs dar",
            taxContent: "Einkommen und Gewinn werden mit dem konfigurierten Satz besteuert",
            costContent: "Gesamtpreis aller Materialien (Anforderungen, Bauplanmaterialien)",
            incomeContent: "Preis des Artikelpreises multipliziert mit der Menge nach Steuern",
            volumeContent: "Durchschnittsvolumen während des Datumsbereichs",
            saleIndexContent: "Dieser Wert soll Artikel hervorheben, die leicht zu verkaufen sind und ein hohes ISK/LP-Verhältnis haben. Er wird berechnet, indem die Punktzahlen dreier Indikatoren multipliziert werden: tägliche Transaktions-ISK-Menge des Artikels, durchschnittliches tägliches Volumen und ISK/LP",
            unitProfitContent: "ISK-Gewinn nach Steuern pro LP-Punkt",
            blueprintContent: "Der Baupreis in der Tabelle und die Bestellung basieren auf der Bauplan-Ausgabe",
            claimContent: "EVE Online and the EVE logo are the registered trademarks of CCP hf. All rights are reserved worldwide" +
                "All other trademarks are the property of their respective owners. EVE Online, the EVE logo, EVE and all associated logos " +
                "and designs are the intellectual property of CCP hf. All artwork, screenshots, characters, vehicles, storylines, " +
                "world facts or other recognizable features of the intellectual property relating to these trademarks are likewise the intellectual property of CCP hf."
        },
        corporation: ' Bitte wählen oder geben Sie einen Konzern ein',
        title: 'Konfiguration',
        data: "Daten",
        quickBar: "Schnellleiste",
        description: "Beschreibung",
        quickbarTitle: "Schnellleiste",
        sourceList: "Alle",
        targetList: "Schnellleiste",
        close: "Schließen",
        reset: "Zurücksetzen",
    },
    table: {
        name: "Artikel",
        quantity: "Menge",
        lpCost: "LP Kosten",
        iskCost: "ISK Kosten",
        materialCost: "Materialkosten",
        price: "Preis",
        income: "Einkommen",
        profit: "Gewinn",
        volume: "Volumen",
        saleIndex: "Verkaufsindex",
        unitProfit: "ISK/LP",
        lookUp: "Suchen",
        order: "Bestellung",
        operation: "Operation",
        error: "Fehler",
        material: {
            lpStoreMaterial: "Austausch Material",
            bluePrintMaterial: "Produktions Material",
            name: "Name",
            quantity: "Menge",
            price: "Preis",
            type: "Typ",
            cost: "Kosten",
            error: "Fehler",
            copy: "Kopieren",
            copySuccess: "Kopiert!",
            copyFail: "Kopieren fehlgeschlagen!",
            sum: "Summe",
        },
        err: {
            productBuy: "Fehler beim Abrufen des Kaufpreises für <b>$1</b>",
            productSell: "Fehler beim Abrufen des Verkaufspreises für <b>$1</b>",
            materialBuy: "Fehler beim Abrufen des Kaufpreises für Produktionsmaterial <b>$1</b>",
            materialSell: "Fehler beim Abrufen des Verkaufspreises für Produktionsmaterial <b>$1</b>",
            requirementBuy: "Fehler beim Abrufen des Verkaufspreises für Anforderung <b>$1</b>",
            requirementSell: "Fehler beim Abrufen des Verkaufspreises für Anforderung <b>$1</b>",
            buyOrder: "Keine Kaufaufträge auf dem Markt gefunden",
            sellOrder: "Keine Verkaufsaufträge auf dem Markt gefunden",
            order: "Keine Aufträge auf dem Markt gefunden"
        },
    },
    calculator: {
        title: "Taschenrechner",
        empty: "Keine Zeile ausgewählt!",
        close: "Schließen",
        calculate: "Berechnen",
        lpCost: "LP-Kosten",
        iskCost: "ISK-Kosten",
        materialCost: "Materialkosten",
        profit: "Maximaler Gewinn",
        unitProfit: "Maximaler ISK/LP",
        reset: "Zurücksetzen",
        materialList: "Materialliste",
    },
    order: {
        sellOrder: "Verkäufer",
        buyOrder: "Käufer",
        history: "Verlauf",
        orderId: "AuftragsId",
        systemName: "Ort",
        volume: "Volumen",
        price: "Preis",
        unitProfit: "ISK/LP",
        expiration: "Ablauf",
        lastUpdated: "Zuletzt aktualisiert",
        statis: {
            lpRange: "ISK/LP Bereich",
            number: "Menge",
            cost: "Kosten",
            income: "Einkommen",
            profit: "Gewinn",
            avgLpPrice: "Avg ISK/LP",
            unitProfit: "ISK/LP",
        },
        stock: {
            average: "Avg(d)",
            minAndmax: "Min/Max(d)",
            average5d: "Moving Avg(5d)",
            average20d: "Moving Avg(20d)",
            minAndmax5d: "Min/Max(5d)",
            volume: "Volumen",
            price: "Preis",
            rangeSelector: {
                month: "1m",
                threeMonths: "3m",
                halfYear: "6m",
                yearToDay: "JbH",
                year: "1J",
                all: "Alle",
            },
        },
    },
}
