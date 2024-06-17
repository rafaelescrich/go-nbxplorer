package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Address represents a Bitcoin address
type Address struct {
	Address string `json:"address"`
}

// Transaction represents a Bitcoin transaction
type Transaction struct {
	TxID      string  `json:"txid"`
	Amount    float64 `json:"amount"`
	Confirmed bool    `json:"confirmed"`
}

// GetAddressesForScheme retrieves addresses for a given derivation scheme
func GetAddressesForScheme(db *sql.DB, derivationScheme string) ([]string, error) {
	rows, err := db.Query("SELECT address FROM addresses WHERE derivation_scheme = $1", derivationScheme)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []string
	for rows.Next() {
		var address string
		if err := rows.Scan(&address); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

// GetTransactionsForScheme retrieves transactions for a given derivation scheme
func GetTransactionsForScheme(db *sql.DB, derivationScheme string) ([]Transaction, error) {
	rows, err := db.Query("SELECT txid, amount, confirmed FROM transactions WHERE derivation_scheme = $1", derivationScheme)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var tx Transaction
		if err := rows.Scan(&tx.TxID, &tx.Amount, &tx.Confirmed); err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}

	return transactions, nil
}

// GetUnusedAddressForScheme retrieves an unused address for a given derivation scheme
func GetUnusedAddressForScheme(db *sql.DB, derivationScheme string) (string, error) {
	var address string
	err := db.QueryRow("SELECT address FROM addresses WHERE derivation_scheme = $1 AND used = false LIMIT 1", derivationScheme).Scan(&address)
	if err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", err
	}

	_, err = db.Exec("UPDATE addresses SET used = true WHERE address = $1", address)
	if err != nil {
		return "", err
	}

	return address, nil
}
