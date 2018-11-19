pragma solidity ^0.4.7;

contract verifier {

    struct verifiedDataBackup {
        uint timestamp;
        uint blockNumberStart;
        uint blockNumberEnd;
        string blockHash;
    }

    bytes32 public verifiedHash;
    verifiedDataBackup public verifiedBackup;

    mapping(address => bytes32) unverifiedHashes;

    function verifier(address[] verifiersList) public {
        // require to not add the same account twice
        for (uint i = 0; i < verifiersList.length; i++) {
            address currentVerifier = verifiersList[i];
            verifiersMap[currentVerifier] = true;
        }
        verifiers = verifiersList;
    }

    address[] public verifiers;
    mapping(address => bool) verifiersMap;

    event VerifiedBackup(uint timestamp, uint blockNumberStart, uint blockNumberEnd, string blockHash);


    modifier onlyVerifiers {
        require(
            verifiersMap[msg.sender] == true,
            "Only verifiers can call this function."
        );
        _;
    }

    function getHashedStruct(verifiedDataBackup unhashed) private returns(bytes32) {
        bytes32 hashed = keccak256(unhashed.timestamp, unhashed.blockNumberStart, unhashed.blockNumberEnd, unhashed.blockHash);

        return hashed;
    }

    function setBackup(uint timestampParm, uint blockNumberStartParm, uint blockNumberEndParm, string blockHashParm) public onlyVerifiers {
        // Check preconditions NotNull and same value is not set multiple times
        verifiedDataBackup memory backup = verifiedDataBackup(timestampParm, blockNumberStartParm, blockNumberEndParm, blockHashParm);
        bytes32 hash = getHashedStruct(backup);
        unverifiedHashes[msg.sender] = hash;
        if(checkHashes()) {
            verifiedHash = hash;
            verifiedBackup = backup;
            emit VerifiedBackup(backup.timestamp, backup.blockNumberStart, backup.blockNumberEnd, backup.blockHash);
        }
    }

    function checkHashes() public returns (bool) {
        bytes32 check = unverifiedHashes[verifiers[0]];
        bool checkValue = true;
        for(uint i=0; i < verifiers.length; i++) {
            if(check != unverifiedHashes[verifiers[i]]) {
                checkValue = false;
            }
        }

        return checkValue;
    }
}
