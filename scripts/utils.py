// Timestamp: 2025-11-01 00:59:32

import hashlib

def hash_password(password: str) -> str:
    return hashlib.sha256(password.encode()).hexdigest()

