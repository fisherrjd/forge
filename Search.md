## Search Implementation

### Searching Maven Central

To begin we will just use exact package names and share that + group id and let users decide which to use

```
curl --request GET \
  --url 'https://search.maven.org/solrsearch/select?q=a%3Aspring-boot-starter-web&rows=20&wt=json'
```

Create a filtered search so we can properly share in order what package is best
