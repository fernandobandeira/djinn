-- Setup PostgreSQL extensions for UUID v7 support
-- Migration: 001_setup_uuid_extension
-- Created: 2025-08-25

-- Enable pgcrypto for UUID generation functions
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Create UUIDv7 generation function
-- UUIDv7 provides timestamp-ordered UUIDs for better database performance
CREATE OR REPLACE FUNCTION uuid_generate_v7()
RETURNS UUID AS $$
DECLARE
    timestamp_part BIGINT;
    random_part BIGINT;
    uuid_bytes BYTEA;
BEGIN
    -- Get current timestamp in milliseconds since epoch
    timestamp_part := (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT;
    
    -- Generate random data for the remaining bytes
    random_part := (RANDOM() * 4294967295)::BIGINT;
    
    -- Construct UUID v7 format: timestamp (48 bits) + version (4 bits) + random (12 bits) + variant (2 bits) + random (62 bits)
    uuid_bytes := 
        SUBSTRING(int8send(timestamp_part) FROM 3 FOR 6) ||  -- 48-bit timestamp
        E'\\x70' ||  -- Version 7 (4 bits) + high random (4 bits)
        SUBSTRING(int4send(random_part) FROM 2 FOR 3) ||     -- Random data
        E'\\x80' ||  -- Variant bits (10) + random (6 bits)
        gen_random_bytes(7);  -- Remaining random bytes
    
    RETURN encode(uuid_bytes, 'hex')::UUID;
END;
$$ LANGUAGE plpgsql VOLATILE;