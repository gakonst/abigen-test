pragma solidity ^0.5.1;


contract ABIGenTest {
    uint a;
    function x() external {
        a = 1;
    }

    function x(uint b) external {
        a = b;
    }
}
