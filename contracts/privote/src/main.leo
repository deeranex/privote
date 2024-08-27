import credits.leo;

// The `program` scope defines the data types, functions, and state associated with the `privote` program.
program privote.aleo {

    // privote voting token
    record PVToken {
        owner: address,
        name: field,        // token name, privote voting token
        symbol: field,      // token symbol, PVT
        amount: u64,        // voting privote, default is 1u64
    }

    // current proposals id, like auto-increament id in mysql  // 2074281269322187893875field => current proposal id
    // use 2074281269322187893875field (means "proposals" in string) as the fixed key here
    // TODO: use array instead of mapping when possible
    // https://vm.aleo.org/api/testnet3/program/${programID}/mapping/proposal_ids/2074281269322187893875field => xU64
    mapping proposal_ids: field => u64;

    // current votes id, like auto-increament id in mysql  // pid => current vote id
    // TODO: use array instead of mapping when possible
    // https://vm.aleo.org/api/testnet3/program/${programID}/mapping/vote_ids/x => xU64
    mapping vote_ids: u64 => u64;


    // proposals will be public for all users, store them into mapping.  // pid => proposal cid
    // https://vm.aleo.org/api/testnet3/program/${programID}/mapping/proposals/x => cid
    mapping proposals: u64 => field;


    // votes will be encoded by drand tlock with proposl expieration time, then the encoded hash will be stored into mapping,
    // so that everyone can retrieve the encoded votes to count after proposal expieration time.
    // TODO: use array instead of mapping when possible
    // pvid (pid append current vote id, e.g. pid-vid) => vote info
    mapping votes: field => field;

    // vote counts       // pid => option counts, append options counts to field, separator: -&&&&&-, e.g. option1-&&&&&-27-&&&&&-option2-&&&&&-15
    // https://vm.aleo.org/api/testnet3/program/${programID}/mapping/counts/x => xU64
    mapping counts: u64 => field;

    // voting privote will be privote, store them into record.
    record privote {
        owner: address,
        privote: u64,             // voting privote, default is 1u64
        withdrawable_block_height: u32, // withdrawable block height
    }

    // get privote voting token
    transition mint_token(amount: u64) -> PVToken {
        return PVToken {
            owner: self.caller,
            name: 7006912922137281764648909451784002770527598field,   // stands for privote Voting Token string
            symbol: 5264980field,   // stands for PVT string
            amount: amount,
        };
    }

    transition mint_privote() -> privote {
        return privote {
            owner: self.caller,
            privote: 1u64,
            withdrawable_block_height: 0u32,    // default can be withdrawed at any time.
        };
    }
}
