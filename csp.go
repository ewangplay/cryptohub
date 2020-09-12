package cryptohub

// Key represents a cryptographic key
type Key interface {
	// Bytes converts this key to its byte representation,
	// if this operation is allowed.
	Bytes() ([]byte, error)

	// Symmetric returns true if this key is a symmetric key,
	// false is this key is asymmetric
	Symmetric() bool

	// Private returns true if this key is a private key,
	// false otherwise.
	Private() bool

	// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
	// This method returns an error in symmetric key schemes.
	PublicKey() (Key, error)
}

// KeyGenOpts contains options for key-generation with a CSP.
type KeyGenOpts interface {
	// Algorithm returns the key generation algorithm identifier (to be used).
	Algorithm() string
}

// KeyGenerator is a interface that provides key generation algorithms
type KeyGenerator interface {
	// KeyGen generates a key using opts.
	KeyGen(opts KeyGenOpts) (k Key, err error)
}

// Signer is a interface that provides signing algorithms
type Signer interface {
	// Sign signs digest using key k.
	//
	// Note that when a signature of a hash of a larger message is needed,
	// the caller is responsible for hashing the larger message and passing
	// the hash (as digest).
	Sign(k Key, digest []byte) (signature []byte, err error)
}

// Verifier is a interface that provides verifying algorithms
type Verifier interface {

	// Verify verifies signature against key k and digest
	Verify(k Key, digest, signature []byte) (valid bool, err error)
}

// CSP is the cryptographic service provider that offers
// the implementation of cryptographic standards and algorithms.
type CSP interface {

	// KeyGen generates a key using opts.
	KeyGen(opts KeyGenOpts) (k Key, err error)

	// Sign signs digest using key k.
	//
	// Note that when a signature of a hash of a larger message is needed,
	// the caller is responsible for hashing the larger message and passing
	// the hash (as digest).
	Sign(k Key, digest []byte) (signature []byte, err error)

	// Verify verifies signature against key k and digest
	Verify(k Key, digest, signature []byte) (valid bool, err error)
}