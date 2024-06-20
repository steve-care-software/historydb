# Programming language.
    Always program code in golang.

# Package names.
Package names should always be written in camelcase.

# Base package.
    The project should always contains packages named: applications, domain and infrastructure at its root.

    The domain objects should be written inside the domain package.

# File names.
    Files should always be named in snakecase.

    When writing an interface, always make sure their names are in camelcase.
    All interfaces should always be written in a file called sdk.go.

    The structs should always be written in other files named by their struct name.

# Base struct writing.
    When writing a struct, always make sure the struct that corresponds to an interface has the same name called, but in lowercase. Always make sure the constructor is named: "create" with the name of the struct after that.

    When writing methods of the struct, name the current instance obj.

    When there is optional members on an interface, always add a method named Has + its name, and make sure its non-nil or true in its body.  It always returns a boolean.


# Domain objects.
    The package name of a domain object should be the name of the domain object.
    A domain object consists of an interface and a struct that implements that interface.

    The domain interface begins by a capital letter.

    The domain struct begins by a lowercase letter.

    A domain object also contain a builder interface and a struct that implements that interface.
    
    For each struct, create a builder interface called by its name + Builder.
    
    It should contain methods called With + name of the property, then pass that property in parameter.  It should contain a method called Create and another one called Now.  Each method of the builder should return an instance of the builder interface, expect the Now method should return an instance of the matched interface, or an error if the mandatory members have not been set.
    Always init the builder member to its default values.

    
# Domain object structs.
    In domain object structs, for optional members, create additional constructors.  In this case, the default constructor should be called create + name of the struct + Internally.  Then create another constructor that calls it called create + name.  Then create another constructor for each optional parameter called create + Name + With + name of the optional parameter.

    The constructor named create + Name + Internally should be written after the other ones.

    When combinations of multiple optional members are mandatory, create a separate constructor named create + Name + With + FirstMember + And + SecondMember.


