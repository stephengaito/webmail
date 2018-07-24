//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"io/ioutil"
)

//
// EmbeddedResource is the structure which is used to record details of
// each embedded resource in your binary.
//
// The resource contains the (original) filename, relative to the input
// directory `implant` was generated with, along with the original size
// and the compressed/encoded data.
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
}

//
// RESOURCES is a simple array containing one entry for each embedded
// resource.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES []EmbeddedResource

//
// Populate our resources
//
func init() {

	var tmp EmbeddedResource

	tmp.Filename = "data/folders.html"
	tmp.Contents = "1f8b08000000000004ffbc586d73db36f27f9f4fb145faff5b1e9ba46ca74d6c93ba4995c449ebc489eda4975eef3a10b9a26083000380b2585fbefb0d1e4851aedba6f7e290c98800f6e1b70f582c9c7ef5ec6c7af9f1ed7358988a4f1ea4f607381565465090c903807481b4b01f00a96186e3e485e4052a3865daa4895f7a0076bb424341d00a33b26478534b6508e4521814262337ac308bacc025cb3172935d6082194679a473ca31db730a01bca0853175849f1ab6ccc8d40b892edb1a07220dae4c62211f43bea04aa3c9de5fbe889e7472be8a2238a506b5815c5635e358001505544cb039c302a617171045013e67e21a14f28c68d372d40b444360a1709e110b461f254945577921e29994461b456b3bc96595f40bc9417c103f4e72add76b71c5449c6b4d800983a562a6cd885ed083278fa2ef3e7c64ece2d50bfc61af38a9be3f7f7adde6cdcba72fcfcb83fdb3ea7d7e73f3588a83f38f45f9e803dd795b5d5cea5f931fbe7db29c15cfaf168f1a02b9925a4bc54a26324285146d251b4d263e26d60557ef1a542d300dcafa536101b3167ac8d07b40e78ad506b4cad716e7b2c0f8ea9395e02cf59fd15ebcb71f3f72965ddd6b98583ea5e39d6fdf958739bff9383d39795b3f7b77faa13e7d23c6f3737a7e76f5f293de3b308f4ecaab83e7bf7e33febb3879f7294fd493d70b71feec2702b9925a4bc5ee1a96261e67089bb5f00f83fc3d5dd20b6fd91f5afaa5b1bdba1bdafb1d70997ff3ea1d9b8df71f7f5ab65717afe72fafce5ed3d3eb79f3e387d54fabf76fc5f4fba78ff97e35fdf1cdabfae4b03a993e7b7273f2e655fef6d9e3cb15255fe0001f6397ae60da1a337f246cae79efb8730db78e0e006634bf2e956c4411e5924b75f4707e68ff1d7b82cffe27d6069718cda534a8d6cc352d0a26ca237834ae57300e3c005e123c3c3cece4dca308363501cca42a504546d647b05faf404bce0a788887f65f907d2f9e58d065648bd2105bc030e334bf0ecc0073294c7483ac5c982398495e841d2f364d9cdfbc9fbad41f38f18a2ea94fb4ce975f8fa0907953a130b01d2ba4453b9a3722374c8ad136dcc2d7a32d26eac61c195c99a325d36cc6f168ce94365bdbf15ce68d1e6d1fc3e76de7e7611ea789b5c8624967b268ed07402ae812724eb5ce88a0cb1955e07fa202e7b4e13d30480bd653da824b994015cd79c38a0ebc9537a00a82ac5654031a8074d618234570859f904d18919165c9d17a9dd35a6341a0a08686e58ce4d2af77cb54956832f2d02b254015a311ae6a2a0a2c3232a75c6358b5e895e4bdc51bd000525d53d181d12a9282b76472e9e108ba6425b5d148134be7bdd88d0d56964b11cda82293ff15699a78570e41a5b433c57b269a292a8aeece49c8e4479c5594f134a16bb63429d87230b541654597217723d585a28fd58643d3860f105807cea88a045d6e5001a49c4d520afe2e4cc8e4a5acd0824a13ced65000d2a4e11bf3df93dfa952f670fe56190cb47159cac690c9a9fbfd73ad1bfe194cd244d0e037e7b370b094bce9f50fd773c9235d4507d6715155440764e2cbad1d4342664f7c64ab6add0bb26333dfd644112d0a29c824651d8492b7f5c26624f45f9146aaf2854d4f764f8aa64e5e38a2065786b81498336e5091d08375b325e50d668440cd698e0bd7bd65fdae2b82a13b3bda1b8fff6f6044f09eab56760ccde64c77567bdd4eae1e7003dcde2a2a4a84d8b78c1a3efbd2ebc73afbd7b22266b0829911f67f57e6c0cb261d566b7044392bc511c779dfa3259e2cb9bd8d3f7f26a024c7acab5f13b76893a737c6c243510c20057303bae1311b1a1e12e3b04b8cc381c9e9627f72b64465dbdf3459ec0f52e6f696cd217eae94544395b5451656d3a45ec3bbbd45ae7193f42d47aa113472cc0dd0e0173012ac42a8506b5aa28eef0a1a5839306bf8594ffe5fcc747d6c39ff9b89735a1a3a065b8dfc675f8dfc14868d4517a175d5ee1db9e96ee1afb37ef7ee75e78ee7b7836db853d736b65c31eb4edeba9520930bd9a81c612a0bbc5bd61ccfba24758f8192994533735db1be5e2537be5893c9d325659cce388214e0896ce6dd95ba592c07d1b063d3057f6a227c998d5359b7aee4dec5f23b16ba78c5d758d5f19c256472611b43f801abfa2f1a34b02e4d7c32849cf12dd6a094ddd37825095ca081a994d70ca1a65a3351c21b5ae1aeaf6ebb50d056031380ab9a2974e9085d6f06b963fc45a319d9e7e9aeab88bb9ef417cbb9dd77c84baa3a724b0a992ba6a17104b7ddf319cc04dec0336a70b4dd930cb6638dc6ed0ed7cab0b60d3b6bda56af050c2038a019ea9cd6387213cb355acb6b759665a2e17c1bfe0684c01190e32055676467a8d7c8f797d30ba398289d6e720c35358b2c21b0d6ddf5b7b17759e67fbc2b768064047620ac3938c1eacfe18e48123847a3182e116820b46f4debcafb6352a2190581966833106c77b5dbeefaedec0eb458d79c991139266bf473a960c4b2f1314b3d53cc51946671cc7676b6bd7e3bfa60dbb10a36fe83fd33d6cd4c1b351a0795768989025767f311c9c8f65a911ded6f19bdd2bb6c3b7b9b8cab6c152b7497f128f9d7cf7ae7df3feb9daf9372970c4cb183cd61b4ca821aef9eb511766c18628742d328018d0829d36e6a86f0f471235cc583d0bd70ed039805869b25b217f3aefb93498d4a336dc02ca8018deecde3c1f4a7ccf71baf29e333b9b2ef21bfef0f8d6f58b2758295689e73b46fa9efda57c568cb736fad01db63e01721f3ecb14b399bc8758d6a4af5f0d85972cef4fd2af477edd4bea36cc9186df93e61a0cae58d15c02083f1313048adac2e7bc0a6cfd0d996d4662b6496ccc75ba07a79f9fa341c083b6cf82cd526e03ea3bc6ddb30c960bc1157ce9c48d798c505d335a72d64b06583e1faa3ad8112db24747eb6e3f7988514b8e60bb1b76f775b525d6584ae25b47f96ab9be0edbe360cf2e41c696e6cd7912f6c7ba7630fc07ae5af86d9d1c7525c63dbd4908580871c3a5e57953301352d31e29216b0a0a2e00854b45d89b1010c196a1fe7deb54c940364b862dad86b230b4caef4f446f776daa8ad890761f7505d0a42d65378d321a80cc0fbdbc0fb79f3a13f93453b7990260b53f1c983ff040000ffff5ef15a9e82150000"
	tmp.Length = 5506
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/login.html"
	tmp.Contents = "1f8b08000000000004ffcc586d73dc28f27fef4fd14bb6fe6bd7fe35f253627b2c4d9dcfceda499cf831c979df2181348c11c8803433d9daef7e0548b23499cd65efea1e70d50850d3fcfad74dd372f4c3d9d5e9fdc3f56b989a824f3622fb008e451e232ad06403209a524c6c072032cc703ab9943913704b9f2aa62889423fbb0156a2a00683c0058d51cde8bc94ca2048a530549818cd1931d398d09aa5347083ff0726986198073ac59cc63b6e4f00af686a4c19d87dea189d7a25c1fdb2a43d95862e4c68511f433ac54a53137fbcff25386cf5fc100470890dd5065259948c5302581028986019a3044eefee20081af89c89475094c7489b25a77a4aa9413055348b9105a3c76158e0454ac42891d268a3706907a92cc26e22dc1bed8d0ec254ebe7b951c1c428d51a011386e68a99658cf414ef1dee077ffdf4c0d8dd9b5fe8bb1d725ebcbd3d795ca6d5c5c9c56dbeb77b557c4ce7f30329f66e1f48beff09ff7c5ddcddeb2fe1bb578775425ecfa6fb15825449ada562391331c2428a65212b8d26de279682d94d45d512980665f9549440b2840e32740ce854b1d28056e9b3c5a92474347bb21a9ca5be1bec8c767647fbceb2d95ac3447d82b77f7e75931fa57cfe707a7e7e5d9edd5c7e2a2f3f88edec16df5ecd2e9ef4ce9ed93fcf677bafbfbcdcfe9b38bf794a4375f87e2a6ecf7e45902aa9b5546cd5b028f4381bb7590bbfe9e4b7b8c677deb26f5afabdbe9dadba763d01f7e9cb37372cd9de3d78aa97b3bbf7d9c5ecea3dbe7cccaacf9f16bf2e3e5e8bd3b727077cb738fdfce14d797e549c9f9e1dcecf3fbc49afcf0eee17187d0701dec72e5cc12c4b1afb236163cdb303ee6cc36f4e1020c1e963ae642548904a2ed5f8457664ff8e9bf7bf37cf9136b4a64126a5a1aa5b5d624298c8c7b0bf5d2e60bb5de435c18ba3a34ecf57fbc0ca46895484aac0c8720cbbe502b4e48cc00b7a64ff5ac5ebd18c04ae039b977ac81a0809c7e963bb3a93c20473caf2a919432239695f78b551e868f334b591dfe370866bece3aca1f2c74d2032ad0a2a0c6c8d14c564b9995522354c8acd2df80d7edcfc8989b232634317665c33cd124ec71953dafcb435ca645ae9cdad63f87dcb91dc8fe228b4f658285122c9d2766c8b04ae21e558eb18095c2758817f048466b8e22d32db22c23a599b723113540519af1869f0fbd6976b94d9bda91a48014449658c140d237e8086600223f39c534b3ec7a5a60401c10637d3314aa59f6fa7b1caa989d10bbf2d02ac180ee8a2c4825012a30c734d9b596b8192bcb37b051c40a44b2c5a385a0552f0259adc7b4002d72cc7d62f5168e55a3edb3658cc522982042b34f9cf0a47a12775082ec2ad519ea520515890f6160ad1e4334d0acc7814e2fec22824ac1e4cd87060a48d9b55cfb5aee97cb7427054f11e0e4b6882552070bd22071071368930f85b3244930b59500b2e0a39eb030288c28af76756300f865128703df1e90d06316b23c3c5760f49672cb7c549221708dcd98e518155ce844b332fb7cbc571c743819915b4f60705095eb51d99659a9a60cf8d75111cb69de6c56e6fdb21b0120bcac1fd064c641241dfd835b2eedc31910f34dae6cc698ebdd7e78a2ce46baf014fbead9bb2fe6f4868d2b663616fbb5c742478e536e3ac826de0360a08d325c7cbb190822217558ee80073ea2a3c9fa1dc08dc6f40b0c8a96ab9dbd9b5876b25426d8b32a98a678576d481b383602a15fb2285c11c81929cc66e1a0176593746a10382a0a0662a498caeafeeeebfe273684c13128934461663d87dd92384d9f41dd80bb25ca36525ebf484034c88146812b1d66f395f96539b5aa0eb05399709b54cb03f4835b6454eade3c456d2a849c0862e6c2ded9976d4d883a024474dc13d95da20a831af688cda5595e2084a8e533a959c501523566057c7dae7882e705172eacaba35f6ae092bdbfaa1f5df23b3d2f6ca8afe2d5c7ae23b2e0704da7d4756e02f8ef1e119b7ed7f9a352ed3c73f118125d67a2e1569e3e979fc8d48b4422b51f7202b0576de29fb5738b305e34e3f87b9b3e04e2c9415e781b2f5de9a2d9a24d0e04e25b7d97d67d77d9eda4a63ed12688fa3af0b759514ccf80cd8f66d28c4dda88919f791dc65b2c408488cf0b7c2d7c6db9b714d6efc83e928b4160fb57c15722b1383613370352840544efe4f24ba3c8ec272f2cf0c9c9ea8f950b00587ef76c6fb21f4bf27da54fe5ca675acb8dcd2f948acbfe9bbf7feeaee16db362c5c06af5cadd22e7efe8640933b59a994c2a924744dd5d2af6fdaff02e4cc4cabc4e54dfdb808e7be264393931a338e134e410af042eb6aa16125d438c4a2b76d48c13f3411becfc653592eddd1f84e0b9dbf468fb42847190bd1e4ce7e0fc23b5a947fd2a09e753674edf7a48d99284c24594e36a2706a0a3ed9f87b000000ffffac1983648c120000"
	tmp.Length = 4748
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/message.html"
	tmp.Contents = "1f8b08000000000004ffbc396b77dbb692dff32be632de95746c924e9cdb2436a93da99d579b3489ed34dbdbdbed81c891049b0218007a5dadfefb9e01f80065354dbae7dcf14944e231ef190c86c9df2ede9d5ffff2fe394ccdac18de4be8070a262669802218de0348a6c8727a00480c37050edfa2d66c82b0d944d563f44216392ad86ee1d7cd069ae18faf2f60bbfd2d89ddc67b40486668180836c33458705c9652990032290c0a93064b9e9b699ae3826718da9723e0821bce8a5067acc0f481650bc0219a1a5386f879ce176970ee9084d7eb123d9406572626c1ce209b32a5d1a41faf5f844f6a3c7f0b4378c30c6a03999c95bcc01c98c861c6051f73cce1fcea0ac2b062bfe0e216141669a0cdba403d4534014c158ed38098d1a7713c63ab2c17d1484aa38d6225bd6472163703f14974123d8e33addbb168c64594691d001706278a9b751ae8293b79f228fcfee75f38bf7afd027f7c90bf9cfd70f9ec769dcd5f3d7b75393979f86ef6315b2e1f4b7172f94b3e79f4333b7c3fbbbad6ff8a7ffceec962943fbf993e9a079029a9b5547cc2451a3021c57a26e73a183a9b900a6e3ecc51ad816b50a44f85398cd6d0b00c8d0674a6786940abac9538933946379f098395d43d860fa2070fa34756b29bbd8289c533767cf8dd87c9d3ac58fe72fef2e5fbf2e2c39b9fcb373f89e3f125bb7c77f3eab37e70621ebd9cdc9c3cffd7df8fff5bbcfcf0398bd593b7537179f18f003225b5968aef0a96c48ecfca6c24e1178dfc035bb02b27d91725fd5adbdeec9a76bf02aeb3bfbffec047c70f1f7f5eac6faede8e5fddbc7bcbdedc8ee79f7e5efd63f5f1bd38ffe1d9e3e2e1ecfcd34fafcb974f672fcf2f9e2c5ffef43a7b7ff1f87ac582af5080b3b1755730eb12531712e46b4e3b36fa6163d701c08865b71325e7220f335948757a7ffc94fecedc82adfb89b4c1058663290daa7673c9f29c8bc9293c3a2e57705ced017098e0fed3a7359e3d84a04b096024558e2a34b23c8587e50ab42c780ef7f129fd55b8f7f21309b6082975f9bc553c8c0a96dd569b01c6529870897c3235a73092455ecd38b4496cf5e6f454bbbea7c41bb660ced16a5d1ef42197d97c86c2c02052c8f2757f3c1799e152f407b081837e8f8b726e4e0daecce9826b3e2af074cc9536bd413496d95cf70767b01d583dfb7e9cc42411f1928c64bea6078044b0056405d33a0d045b8c9802f713e63866f3a2610c929c372b29e1322e5085e362cef39a79c2e7adaa10115554de1a80643437468a4a15ee25e8b2111a399914485a2f58a9310f20678655c3699049375e0f3335419306f71dd10098e22cc455c9448e791a8c59a1b11a25ee952c1a893bac0124ba64a26646ab508a621d0caf1d3b822df88491359298d6392dd6d0d9ca3329c21153c1f0dfb534899d2a7da612568be234138e1413797de6c4c1f0138e668c1749ccda6d499cf385f74a46e579ed21bb96aa4dd1d8aaa3d0645e781c9002474c85822d3aab0092820f1306ee2c8c83e12b3943622a890bdeb20290c4f3a2f3fe47f86b528a82f32e31f0a8157222e72618beb1bf7f4eb5a31fef258905abf4667556059692cb86be3f9ec922d4b3f0841417cef2f0a459d50d374e011f52522dbd153b9eea2d0a599e4b110c135e733029d6e5941c129aa7502353d994bc93eff1d0c4e2ab22d4e0ca04d603c6bc30a882aa04abdf16ac98631a0450162cc3a9ade7d266d6e6c0aa383b7d707cfc1f9e109ef2087ced145cd7423bd216adf636036c368a89094255436ad8bac4eba0f5fd1657c80dce606404fdab931c38dc41cd2ac91bb2824fc46981e3a6428bddb278b389b6db00942c30adb3d7d00e92ebd8c4eb60b341917b2c75a4edbe508855ee52b9c5d3da2d9e7a226f367c0cd173a5a4f2f19644be1a4de2b2e561b3c142a3bf946a9957c87254baa9550812c346946fdd59e05eecffa1368a9798576f53b940553dbb03b699caa4c85150aef61970163ab8c5f5111c584f81d3b42df36b5e3c1ee9b2a08689c9879b8ddd08db6d129bdc0e018d3934f5686c544b909c6257eb568e7689d5c133635836a583b6ab87cd0648c5f5e5e415d3fe4a8f4bdf644b2c0a98b9ab4dc85acc9e2288af4a15fc080eba4af0487814ba59c0cf220e7c0e2aa7f9ae4391a0732635d1efe5819295a8b28297d5f138e5798e220d8c9a236587bbc71c897280d10b5ee04f6ce67b9783fe86a6ab0b15dda7b6db81171594c33be70bc1d7c9d21c667500978acf985adb605ee9a043c5422748ef4ed775c3efa38289db3d0baaf3c8d934de6c0e1acf70d7d6edb633f8f1f5851bf1d5b307ad2d4f0a36a2abe0855c8a4232bf8caae14f8d97d77b1dc6afb11d404d70471d3bc96b8f993a298be08b098edcc4456243c846def7325f77534f736a3ab87ba487868dbac144c87703f5d5f5db377e841224457308de8dec6a03cb0c5f6025cbdd00aa8b92fb543687d40ba8b4dd1692de8c7338c34675656a0bc7d48e0c89e01e35772b1c2b1a8afcaf8852e7fb6f12898ebbfd22b9992f8a748d2b5b31353626d82dda5c79d75902743cbbe8728a556cf99799b864cb3f556bb770f44a35023ffd1836b2a645d12d1bed9a461725135834a530eda11118b30c818b3dc708d91eb68d652a0bbbdaa6f51f3af29a14536db913765fc949fe654e88472c3452e7adf697bb5c59171826a542628dbcbfe58f6c4f87304dfebf98f4b4406ef007e42e9f7db267fe1e723bd4ab58ae1cce9bac1eed44520eff538c74794615d35f797158aa1e06dd8fdc63e314ee15fc56475d35b6f7c826df58a356e51765167bc16e66bbb5009df6b37ce7b4efa64d6f27819707dbe64630bc927395219ccb1cf7c76c93fdeaf6e4849be97c64fb74fa76152fddf531183e5b305ed812520a708b282477b1ee4661a710e8aac0de89ba054d5744f83a19cf65b9b697c05d5e5c56ba23a1b557748bb3321af338185e51ab0a7ec459f98d0255ce66d51f3b67a87cc6357dbcdbd5de5650dbf869ba6b07fd5e64ac968dfad536442885fed61b44c8b2e9be1d00077d33e57a40ede17e2f9b2b2d55efa8574aea14abde20b2157ddfb259c11ec20e6a542ccfcfe9a2d0efb9c4d11b547d2f07dba36fc2a6702617f8458483488a7e6f26e71ae765efa8e110fa3868db870ef4929b6c0a7d8c96539e4dbba5e7ee62803886373836705ef0ec36f21927c898467870ba3b5c37e9a24266b61d0469ab1c6354bfd718674737042385ecf6cc65210fe218def23cb7b7b03fe4e5e11d5e16d4b4c3e5272e72b9fc263e96764b244b14fd06c311f45c35bc87718c4a850b14e6c27507fb5dbb1354b2d163037522b6d0de05b683b37b9512e89166fd862535385c4ffd8b2dd338862b34702ee52d472899d65c4c802e2647ae317104395b6b3a94715572859650eb4199ddf8bb46d3a70f4b47f66e79e496fe4e3bdbe8234d57cb6929a4b60fd2b83e4d37fb0ca6029770c10c7a5af2a6238dc6cefa63936a6c0087eddab56ed54c342a162ca329ea8c95d8b72fb4abdfe25beb344dc5bc2806f05f1004700ac1598555a7c1a14fd7c88fd7e757467131b1b48333289999a671002dedc6e91d03a9fb71aa3884200de0b0c35ba5986d65e338864b348ae30281550be92b11a972bf4d2668fa1e91ae21f8d1ea687de4a6d31dd6225d16dcf483b3a0e57e2c15f4797a7cc613b7292a504cccf48c1f1e7a39a249b504ab4ac65ff96f919e8fb451fde38a240d7191e3eaddb81fa4c1a02544b0bebbd111dddd76f8a0bb7195ae2285b68fd68fffe79ffaf07fffa90f0fe2c951e08942c0c7d05fa515196b034f08828e20040acd5c09988bca65d65dcae0876915b19ee95ed8ce1f982956ddb2909a6a47f66367894a736dc04c99018df66bc58e455dabf02de3c548aee84b869b07ebd0aed798b60e3641f3bc406ace7cbf7e9df77b6eb79790280cdc20a46e7b6423801cb92c519d33ed871d2d2fb8de4f427fbfb6670fa58c7ecff5f83c52d66f080187148ecf804342b86aef01721f5fd9b4945c1a525ae6ec2d50512d5f050401998f5675196e3ccac93680610ac71dbb16dca2b43dd528e7ba2cd81a52e891316c6fb3e711a1eabed633c11f6d165260bbafb23d7d75a3946a3323d4dd5cfaa05ece2b6d37b9c1f3934b64990123e973b998a0ae4e54d2cab79ad9ae8fa4b8c5f5bc84b43278e543f5f119c7f04e40c92618521305a64cd009cac4ba4e3164c0ca43e9b39a532d17138f335c716de8d848ab4d36f53442377292d5dac59ed91dabd605216d56d4ba77242bc69bd3c0e9d93ff1927824f3f5f05e124fcdac18defbbf000000ffff116de7a562210000"
	tmp.Length = 8546
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/messages.html"
	tmp.Contents = "1f8b08000000000004ffec396b73d4b896dff915e71af6767725b613c21d4862f71693406086671260b98f9d52dba76d256ec948723f6e36ff7deb48b2dbdd04ee305bb5b51f5614694b3e3aef97e4e44fa76f4f2e3fbf7b06a59955e37b09fd40c54491062882f13d80a44496d3034062b8a9707c7303d17359e5a8e0f63689dde23d2080191a0682cd300de61c17b55426804c0a83c2a4c182e7a64c739cf30c433bd9052eb8e1ac0a75c62a4cf72d490087a834a60ef14bc3e76970e2908497ab1a7b280d2e4d4c4c1f435632a5d1a41f2e9f874f5a3c7f0a4378c50c6a03999cd5bcc21c98c861c6059f72cce1e4e202c2d0b35f71710d0aab34d06655a12e114d00a5c2691a1033fa288e676c99e5229a4869b451aca649266771b7101f4407d1e338d37abd16cdb88832ad03e0c260a1b859a5812ed9c19347e1cf1f3f737ef1f239feba9f9fcd7e397f7abdca9a174f5f9c17070fdfce3e648bc563290ece3fe7c5a38f6ce7ddece252ff33fef5a727f349feecaa7cd4049029a9b554bce0220d98906235938d0ec6ce26a482abf70daa15700d8af4a93087c90a3a96a1d380ce14af0d6895ad25ce648ed1d517c26025758fe17eb4ff307a6425bbba5330317fcaf6767e7a5f1c66d5e2f3c9d9d9bbfaf4fdab8ff5ab37626f7acecedf5ebdf8a2f70fcca3b3e2eae0d93fffb2f71fe2ecfd972c564f5e97e2fcf4af01644a6a2d15df162c891d9fde6c24e1778dfc0b9bb30b27d97725fdbdb6bdda36eddd0ab8ccfef2f23d9fec3d7cfc65bebaba783d7d71f5f6357b753d6d3e7d5cfe75f9e19d38f9e5e9e3eae1ece4d39b97f5d9e1ecece4f4c9e2eccdcbecdde9e3cb250b7e87029c8dadbb8259d598ba90205f73dab1910d37160e00262cbb2e946c441e66b292eae8fef490fe1d3b805bf7136983730ca7521a54ebcd35cb732e8a2378b4572f61cfef017098e0fee1618be70e42b049096022558e2a34b23e8287f512b4ac780ef7f190fe79dc77f21309360f292df579f33c4c2a965dfbcd0053294cb8405e94e60826b2cafd1b873689adde9c9e5ad7ef29f18acd9973b456970f8690cbac99a130308a14b27c359c3622335c8ae1086ee0c170c045dd9823834b7334e79a4f2a3c9a72a5cd60144d65d6e8e1e8186e4756cf7d3f4e629288784926325fd1034022d81cb28a699d0682cd274c81fb09739cb2a6ea188324e71d24255cc605aa705a353c6f99277c3d288f88a8a2eac1002493c61829bc2adc24d8642334b2282a24ad57acd698079033c3fc721a64d2adb7cb4c1568d2e0be231a00539c85b8ac99c8314f8329ab34fa55e25ec9aa937883358044d74cb4cc68154a51ad82f1a56347b0392f1859238909ce69b11d1b5b79264538612a18ff6f8126b153659fa984b5a238cd8413c544ded69c38187fc2c98cf12a89d97a5b12e77cde9b925179de7ac8b6a55a5374b6da5068d2543d0e488113a642c1e61b500049c5c70903570be360fc42ce90984ae28aaf590148e2a6da987f0b7f4b4a51707e4d0c7ad42a59c8c604e357f6f75f53ddd04f6f92c48279bd599df9c05272d1d1efaf67b20af52c3c20c585b33c3ce8a036c38d53c0879454eb1ec496a7f6804296e75204e384b71c14d5aa2ec921a17b0a35329595e49dfc0e0f4d2c3e1fa1069726b01e30e5954115f816ac9dcd59d5601a0450572cc3d2b66e69f7d6e640df9c1dedefedfd5b4f08af3c9bac68f4d55371dd4aed685bbcbab71be0e646315160db2f6ab87599d78db5f3af7185dce00c2646d0ff36cb81c31db4bc92c021ab78218e2a9c762d5aecc0e29b9be8f63600252b4cdbf435b68be43b9d30c41e8abcc79217d773d78fb2bee0de2f0e5bbf38ec897c73c3a7103d534aaa3ede9ac8fbd524aed73cdcdc60a5b10f4a19df66c78a4da8137dc70aa458f179ad470bfac15bb3820b9bf936205ce07a89dab18ee34e61ad85e0f636a6fefe9dc239901237585138e7b2eb2cfbc3a5560b5cf23c4791064635188cff5cb12f8d3cbe2bc752b2e865351adbd9e48fb2ff0697e62bf669714b3934becdbafa9fb0be99087deae90c911836a1f2e94abb9bd8bfa1368ad798fb5929e7a8fcb3eb97ba579914390a2abd7dfb26468d13538e9398fe98727cca0c7693e74acebac94533b9c2ccb8796c541f4d17b8af516b56e076e41ae5eabdaf0733074591f7f2d4bad0037f44241fbab9018a8a37b880db5b9f6ebeeecc020a060ac77162f2b1dff3d4189695d46f1107ce545e695daaec25cd9ad5a8b28ad7be97d8f445ef84372eea3d8573ac2b3a08defe6becba64aaed52be8739894dde0a1191fe09797f8dccb0bde6add12d7f6510ab9bb503c5d65dd636f352ad01fe3f9174baa0410dccc602fc9e3c4839e3ff62222189287f768da0af56f645528fff2c26ba3ea652f347260e8b3ffd5167e91ebbced24da17f486ccbedba03efd2d266ed14ee68d2bdddeca5a8b6cef2f0a7deeb8d32776763da6684f5b130185fc84665082732c7bb8b4a5704db8b9d829bb299d81b0e7dbd8c17aef10ec64fe78c57365b4b010e88da886dacdbf9beb30c8d4d15d86ef27b22c286f8b608de25e389ac57b67ddee6c5eef85a426bafe81a677534e57130bea0433efc8ab3fa0705f2ce66258b9d33789f71c7e55e5f7ae7217a7d64eeee251e0c0791cd6860d4dfec51924acb3f06a3085956deb503e0c1d0945c8fe8626d38c81aa5a51aec0e6a49776c6a308a6cf11c5a36fdb883b01b2d2a96e7275493870396193ec7c1c8df18b871bbfb43d814cee41cbf8b701449311ccc64a3b1a907bb1d8730c4d1fae2c50dbde0262b6188d1a2e45939eaf302dbc000710caf706ae0a4e2d975b4090c90318db07fb4bddc5e6f4495cc6c3b09e95a39c6a8e1a033ce966e684c14b2eb6397857a238ee135cf73dbf07c9397875ff132a7eb0e5c7ce222978b1fe26361b744b24631ec30ecc2e0b749c5c4f51d8c63542b9ca330a7ee5e65b869771a5e367aecc6badc02b4973af6e9f89e57c2ad47d4bfeaa1a3a1bb8dfcee65531cc3051a3891f29a23d44c6b2e0a78c366b8eb8e74bb90b395062e00973557e8786b7d1c32bbf1378d664857f2bbf618b8eb407fa39debe8234d7b700285d49e203bd7a7d7dd3e83a9c005505fd3d352ef75a4d1d8b7fdb5c2af8d60670dbbd25e3b00968667c1329aa2ce588d433ba15dc335be954ed354345535827f87208023088e3d569d063b7dba467eb83cb9308a8bc2d20e8ea166a64ce300d6b43ba7770ca4eec7a962078234809d0ddebc626ebd8de318ced1288e7304e601e97e9d5479b74d0a34c31e914d43f0dde5ee6ad7bd4eb7588b745d71330c8e8335f753a960c8d3bd639eb84d5185a230e531dfd9e9e5882ed5d2587a19ffc6ff11e966a28d1aee7992b4c4458ecbb7d3619006a335211aabaf373aa2dbdb76f637372ed365a4d0de400ce3fffcbbdef9afbfeb9d0771b11bf444a1c1a7305ca69e8cb5414f081a1b82d050681a25a011de65569b94a11fa63e627ba67b6eef4cc094e8af1942ba8dd8b59f896a549a6b03a6640634da7bde2d8bba4b96d78c5713b91cf67236058dbba549d70e56a07956219d677e5ebdcc8703b7bb979068975b84d46d8f6c049023d735aa13a6fb6147e015d77793d03faf6ceda194311cb8b37e8f94f51b42c02185bd63e09010aed67b80dca7af6c0225978694c09cbd05aa1797af5ff980a041e623a84d863b8f72b28d609cc2de865d2b6e51dae36194735d576c05290cc818f65268d0238295f691e5c6b7360b2970bdcfdb9ebe57504ab59911da7b30fa1459375edb5d6ee8f9c939b2cc8091f4a15114a87d4525adfca8992d7c24c535ae9a1a526f70ef436df98c63782ba0660586956439944c50056562d5a61832a0f750fa20e154cb45d1e30c975c1b2a1ba9df64534f27742727596d0ddc33bb63d5ba20a41d848b01f0243de35d35707aee57bc249ec87c35be97c4a59955e37bff1d0000ffff318c60f7781e0000"
	tmp.Length = 7800
	RESOURCES = append(RESOURCES, tmp)

}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	for _, entry := range RESOURCES {
		//
		// We found the file contents.
		//
		if entry.Filename == path {
			var raw bytes.Buffer

			// Decode the data.
			in, err := hex.DecodeString(entry.Contents)
			if err != nil {
				return nil, err
			}

			// Gunzip the data to the client
			gr, err := gzip.NewReader(bytes.NewBuffer(in))
			defer gr.Close()
			data, err := ioutil.ReadAll(gr)
			if err != nil {
				return nil, err
			}
			_, err = raw.Write(data)
			if err != nil {
				return nil, err
			}

			// Return it.
			return raw.Bytes(), nil
		}
	}
	return nil, errors.New("Failed to find resource")
}

//
// Return the available resources.
//
func getResources() []EmbeddedResource {
	return RESOURCES
}
