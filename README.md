# prettyrlp

CLI utility to pretty print RLP encoded data.

## Usage

`prettyrlp cc86616e6472656984616c6578`

```bash
List {
    String andrei
    String alex
}
```

## Building

```bash
make build
```

## Testing

```bash
make test
```

## References

* RLP specs: https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp/

* RLP encoder utility: https://github.com/SamuelHaidu/simple-rlp

```python
>>> rlp.encode(["andrei", "alex"]).hex()
'cc86616e6472656984616c6578'
```
