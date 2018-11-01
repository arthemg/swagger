package dataparser

import (
	"bytes"
	"encoding/json"
	"github.com/arthemg/dataParser/models"
	"github.com/arthemg/dataParser/restapi/operations"
	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


type Message struct {}

func TestGetJSON(t *testing.T) {
	url := "https://api.github.com/repositories"
	var messages []Message
	err := getJSON(url, &messages)
	if err != nil {
		t.Error("TestGetJSON failed SHOULD NOT HAVE error, got", err)
	}

	brokenURL := "THIS DOES NOT EXISTS AS URL"
	messages = []Message{}
	err = getJSON(brokenURL, &messages)

	if err == nil {
		t.Error("TestGetJSON failed should HAVE error, got", err)
	}
}

func TestCheckStatusCode(t *testing.T) {
	var statusOk = "https://httpstat.us/200"
	returnStatusCode, err := checkStatusCode(statusOk)
	if err != nil || returnStatusCode != 200 {
		t.Error("Should have received Status code 200, got", returnStatusCode)
	}
	var brokenURl = "https://"
	_, err = checkStatusCode(brokenURl)
	if err == nil {
		t.Error("Should have received an error, got", err)
	}
}

//func TestPingCheck(t *testing.T) {
//	status := pingCheck("google.com")
//	if !status {
//		t.Error("Testing Incorrect Server should be true, got ", status)
//	}
//	wrongStatus := pingCheck("https://google.com")
//	if wrongStatus {
//		t.Error("Testing Incorrect Server should be false, got ", wrongStatus)
//	}
//}

//func TestJSONGet(t *testing.T) {
//	req, err := http.NewRequest("GET", "", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	testData := &DataURLs{
//		DataLocation: "https://api.github.com/repositories",
//		URLToPing:    "api.github.com",
//	}
//	params := operations.JSONGetParams{HTTPRequest: req, Jsonrepo: []string{"https://api.github.com/repositories", "api.github.com"}}
//	r := JSONGet(testData)
//	w := httptest.NewRecorder()
//	r(params).WriteResponse(w, runtime.JSONProducer())
//	if w.Code != 200 {
//		t.Error("Should receive Status Code 200, got", w.Code)
//	}
//	IncorrectParams := operations.JSONGetParams{HTTPRequest: req, Jsonrepo: []string{"https://api.github.com/repositor", "api.github.com"}}
//	r = JSONGet(testData)
//	w = httptest.NewRecorder()
//	r(IncorrectParams).WriteResponse(w, runtime.JSONProducer())
//	if w.Code != 404 {
//		t.Error("Should receive Status Code 404, got", w.Code)
//	}
//
//	IncorrectURLPing := operations.JSONGetParams{HTTPRequest: req, Jsonrepo: []string{"https://api.github.com/repositories", "https://httpstat.us/404"}}
//	r = JSONGet(testData)
//	w = httptest.NewRecorder()
//	r(IncorrectURLPing).WriteResponse(w, runtime.JSONProducer())
//	if w.Code != 404 {
//		t.Error("Should receive Status Code 404, got", w.Code)
//	}
//}

//func TestJSONGetMock(t *testing.T) {
//
//	assert := assert.New(t)
//	//testRepos := &models.JsonrepoItems0{}
//	tests := []struct {
//		description        string
//		reposClient        *models.Jsonrepo
//		url                string
//		expectedStatusCode int
//		expectedBody       string
//		jsonString         string
//	}{
//		{
//			description:        "Internal Server Error 500 ",
//			reposClient:        &models.Jsonrepo{},
//			url:                "https://httpstat.us/500",
//			expectedStatusCode: 500,
//			expectedBody:       "{\"code\":500,\"message\":\"INTERNAL_SERVER_ERROR\"}\n",
//			jsonString:         ``,
//		}, {
//			description:        "Resource Not Found",
//			reposClient:        &models.Jsonrepo{},
//			url:                "https://httpstat.us/404",
//			expectedStatusCode: 404,
//			expectedBody:       "{\"code\":404,\"message\":\"RESOURCE_NOT_FOUND\"}\n",
//			jsonString:         ``,
//		}, {
//			description:        "No Data Available",
//			reposClient:        &models.Jsonrepo{},
//			url:                "https://httpstat.us/200",
//			expectedStatusCode: 404,
//			expectedBody:       "{\"code\":404,\"message\":\"NO_DATA_AVAILABLE\"}\n",
//			jsonString:         ``,
//		}, {
//			description:        "Get Valid Json",
//			reposClient:        &models.Jsonrepo{},
//			url:                "https://api.github.com/repositories",
//			expectedStatusCode: 200,
//			expectedBody:       "[{\"full_name\":\"mojombo/grit\",\"html_url\":\"https://github.com/mojombo/grit\",\"id\":1,\"name\":\"grit\",\"url\":\"https://api.github.com/repos/mojombo/grit\"},{\"full_name\":\"wycats/merb-core\",\"html_url\":\"https://github.com/wycats/merb-core\",\"id\":26,\"name\":\"merb-core\",\"url\":\"https://api.github.com/repos/wycats/merb-core\"},{\"full_name\":\"rubinius/rubinius\",\"html_url\":\"https://github.com/rubinius/rubinius\",\"id\":27,\"name\":\"rubinius\",\"url\":\"https://api.github.com/repos/rubinius/rubinius\"},{\"full_name\":\"mojombo/god\",\"html_url\":\"https://github.com/mojombo/god\",\"id\":28,\"name\":\"god\",\"url\":\"https://api.github.com/repos/mojombo/god\"},{\"full_name\":\"vanpelt/jsawesome\",\"html_url\":\"https://github.com/vanpelt/jsawesome\",\"id\":29,\"name\":\"jsawesome\",\"url\":\"https://api.github.com/repos/vanpelt/jsawesome\"},{\"full_name\":\"wycats/jspec\",\"html_url\":\"https://github.com/wycats/jspec\",\"id\":31,\"name\":\"jspec\",\"url\":\"https://api.github.com/repos/wycats/jspec\"},{\"full_name\":\"defunkt/exception_logger\",\"html_url\":\"https://github.com/defunkt/exception_logger\",\"id\":35,\"name\":\"exception_logger\",\"url\":\"https://api.github.com/repos/defunkt/exception_logger\"},{\"full_name\":\"defunkt/ambition\",\"html_url\":\"https://github.com/defunkt/ambition\",\"id\":36,\"name\":\"ambition\",\"url\":\"https://api.github.com/repos/defunkt/ambition\"},{\"full_name\":\"technoweenie/restful-authentication\",\"html_url\":\"https://github.com/technoweenie/restful-authentication\",\"id\":42,\"name\":\"restful-authentication\",\"url\":\"https://api.github.com/repos/technoweenie/restful-authentication\"},{\"full_name\":\"technoweenie/attachment_fu\",\"html_url\":\"https://github.com/technoweenie/attachment_fu\",\"id\":43,\"name\":\"attachment_fu\",\"url\":\"https://api.github.com/repos/technoweenie/attachment_fu\"},{\"full_name\":\"Caged/microsis\",\"html_url\":\"https://github.com/Caged/microsis\",\"id\":48,\"name\":\"microsis\",\"url\":\"https://api.github.com/repos/Caged/microsis\"},{\"full_name\":\"anotherjesse/s3\",\"html_url\":\"https://github.com/anotherjesse/s3\",\"id\":52,\"name\":\"s3\",\"url\":\"https://api.github.com/repos/anotherjesse/s3\"},{\"full_name\":\"anotherjesse/taboo\",\"html_url\":\"https://github.com/anotherjesse/taboo\",\"id\":53,\"name\":\"taboo\",\"url\":\"https://api.github.com/repos/anotherjesse/taboo\"},{\"full_name\":\"anotherjesse/foxtracs\",\"html_url\":\"https://github.com/anotherjesse/foxtracs\",\"id\":54,\"name\":\"foxtracs\",\"url\":\"https://api.github.com/repos/anotherjesse/foxtracs\"},{\"full_name\":\"anotherjesse/fotomatic\",\"html_url\":\"https://github.com/anotherjesse/fotomatic\",\"id\":56,\"name\":\"fotomatic\",\"url\":\"https://api.github.com/repos/anotherjesse/fotomatic\"},{\"full_name\":\"mojombo/glowstick\",\"html_url\":\"https://github.com/mojombo/glowstick\",\"id\":61,\"name\":\"glowstick\",\"url\":\"https://api.github.com/repos/mojombo/glowstick\"},{\"full_name\":\"defunkt/starling\",\"html_url\":\"https://github.com/defunkt/starling\",\"id\":63,\"name\":\"starling\",\"url\":\"https://api.github.com/repos/defunkt/starling\"},{\"full_name\":\"wycats/merb-more\",\"html_url\":\"https://github.com/wycats/merb-more\",\"id\":65,\"name\":\"merb-more\",\"url\":\"https://api.github.com/repos/wycats/merb-more\"},{\"full_name\":\"macournoyer/thin\",\"html_url\":\"https://github.com/macournoyer/thin\",\"id\":68,\"name\":\"thin\",\"url\":\"https://api.github.com/repos/macournoyer/thin\"},{\"full_name\":\"jamesgolick/resource_controller\",\"html_url\":\"https://github.com/jamesgolick/resource_controller\",\"id\":71,\"name\":\"resource_controller\",\"url\":\"https://api.github.com/repos/jamesgolick/resource_controller\"},{\"full_name\":\"jamesgolick/markaby\",\"html_url\":\"https://github.com/jamesgolick/markaby\",\"id\":73,\"name\":\"markaby\",\"url\":\"https://api.github.com/repos/jamesgolick/markaby\"},{\"full_name\":\"jamesgolick/enum_field\",\"html_url\":\"https://github.com/jamesgolick/enum_field\",\"id\":74,\"name\":\"enum_field\",\"url\":\"https://api.github.com/repos/jamesgolick/enum_field\"},{\"full_name\":\"defunkt/subtlety\",\"html_url\":\"https://github.com/defunkt/subtlety\",\"id\":75,\"name\":\"subtlety\",\"url\":\"https://api.github.com/repos/defunkt/subtlety\"},{\"full_name\":\"defunkt/zippy\",\"html_url\":\"https://github.com/defunkt/zippy\",\"id\":92,\"name\":\"zippy\",\"url\":\"https://api.github.com/repos/defunkt/zippy\"},{\"full_name\":\"defunkt/cache_fu\",\"html_url\":\"https://github.com/defunkt/cache_fu\",\"id\":93,\"name\":\"cache_fu\",\"url\":\"https://api.github.com/repos/defunkt/cache_fu\"},{\"full_name\":\"KirinDave/phosphor\",\"html_url\":\"https://github.com/KirinDave/phosphor\",\"id\":95,\"name\":\"phosphor\",\"url\":\"https://api.github.com/repos/KirinDave/phosphor\"},{\"full_name\":\"bmizerany/sinatra\",\"html_url\":\"https://github.com/bmizerany/sinatra\",\"id\":98,\"name\":\"sinatra\",\"url\":\"https://api.github.com/repos/bmizerany/sinatra\"},{\"full_name\":\"jnewland/gsa-prototype\",\"html_url\":\"https://github.com/jnewland/gsa-prototype\",\"id\":102,\"name\":\"gsa-prototype\",\"url\":\"https://api.github.com/repos/jnewland/gsa-prototype\"},{\"full_name\":\"technoweenie/duplikate\",\"html_url\":\"https://github.com/technoweenie/duplikate\",\"id\":105,\"name\":\"duplikate\",\"url\":\"https://api.github.com/repos/technoweenie/duplikate\"},{\"full_name\":\"jnewland/lazy_record\",\"html_url\":\"https://github.com/jnewland/lazy_record\",\"id\":118,\"name\":\"lazy_record\",\"url\":\"https://api.github.com/repos/jnewland/lazy_record\"},{\"full_name\":\"jnewland/gsa-feeds\",\"html_url\":\"https://github.com/jnewland/gsa-feeds\",\"id\":119,\"name\":\"gsa-feeds\",\"url\":\"https://api.github.com/repos/jnewland/gsa-feeds\"},{\"full_name\":\"jnewland/votigoto\",\"html_url\":\"https://github.com/jnewland/votigoto\",\"id\":120,\"name\":\"votigoto\",\"url\":\"https://api.github.com/repos/jnewland/votigoto\"},{\"full_name\":\"defunkt/mofo\",\"html_url\":\"https://github.com/defunkt/mofo\",\"id\":127,\"name\":\"mofo\",\"url\":\"https://api.github.com/repos/defunkt/mofo\"},{\"full_name\":\"jnewland/xhtmlize\",\"html_url\":\"https://github.com/jnewland/xhtmlize\",\"id\":129,\"name\":\"xhtmlize\",\"url\":\"https://api.github.com/repos/jnewland/xhtmlize\"},{\"full_name\":\"ruby-git/ruby-git\",\"html_url\":\"https://github.com/ruby-git/ruby-git\",\"id\":130,\"name\":\"ruby-git\",\"url\":\"https://api.github.com/repos/ruby-git/ruby-git\"},{\"full_name\":\"ezmobius/bmhsearch\",\"html_url\":\"https://github.com/ezmobius/bmhsearch\",\"id\":131,\"name\":\"bmhsearch\",\"url\":\"https://api.github.com/repos/ezmobius/bmhsearch\"},{\"full_name\":\"uggedal/mofo\",\"html_url\":\"https://github.com/uggedal/mofo\",\"id\":137,\"name\":\"mofo\",\"url\":\"https://api.github.com/repos/uggedal/mofo\"},{\"full_name\":\"mmower/simply_versioned\",\"html_url\":\"https://github.com/mmower/simply_versioned\",\"id\":139,\"name\":\"simply_versioned\",\"url\":\"https://api.github.com/repos/mmower/simply_versioned\"},{\"full_name\":\"abhay/gchart\",\"html_url\":\"https://github.com/abhay/gchart\",\"id\":140,\"name\":\"gchart\",\"url\":\"https://api.github.com/repos/abhay/gchart\"},{\"full_name\":\"benburkert/schemr\",\"html_url\":\"https://github.com/benburkert/schemr\",\"id\":141,\"name\":\"schemr\",\"url\":\"https://api.github.com/repos/benburkert/schemr\"},{\"full_name\":\"abhay/calais\",\"html_url\":\"https://github.com/abhay/calais\",\"id\":142,\"name\":\"calais\",\"url\":\"https://api.github.com/repos/abhay/calais\"},{\"full_name\":\"mojombo/chronic\",\"html_url\":\"https://github.com/mojombo/chronic\",\"id\":144,\"name\":\"chronic\",\"url\":\"https://api.github.com/repos/mojombo/chronic\"},{\"full_name\":\"sr/git-wiki\",\"html_url\":\"https://github.com/sr/git-wiki\",\"id\":165,\"name\":\"git-wiki\",\"url\":\"https://api.github.com/repos/sr/git-wiki\"},{\"full_name\":\"queso/signal-wiki\",\"html_url\":\"https://github.com/queso/signal-wiki\",\"id\":177,\"name\":\"signal-wiki\",\"url\":\"https://api.github.com/repos/queso/signal-wiki\"},{\"full_name\":\"drnic/ruby-on-rails-tmbundle\",\"html_url\":\"https://github.com/drnic/ruby-on-rails-tmbundle\",\"id\":179,\"name\":\"ruby-on-rails-tmbundle\",\"url\":\"https://api.github.com/repos/drnic/ruby-on-rails-tmbundle\"},{\"full_name\":\"danwrong/low-pro-for-jquery\",\"html_url\":\"https://github.com/danwrong/low-pro-for-jquery\",\"id\":185,\"name\":\"low-pro-for-jquery\",\"url\":\"https://api.github.com/repos/danwrong/low-pro-for-jquery\"},{\"full_name\":\"wayneeseguin/merb-core\",\"html_url\":\"https://github.com/wayneeseguin/merb-core\",\"id\":186,\"name\":\"merb-core\",\"url\":\"https://api.github.com/repos/wayneeseguin/merb-core\"},{\"full_name\":\"sr/dst\",\"html_url\":\"https://github.com/sr/dst\",\"id\":190,\"name\":\"dst\",\"url\":\"https://api.github.com/repos/sr/dst\"},{\"full_name\":\"mojombo/yaws\",\"html_url\":\"https://github.com/mojombo/yaws\",\"id\":191,\"name\":\"yaws\",\"url\":\"https://api.github.com/repos/mojombo/yaws\"},{\"full_name\":\"KirinDave/yaws\",\"html_url\":\"https://github.com/KirinDave/yaws\",\"id\":192,\"name\":\"yaws\",\"url\":\"https://api.github.com/repos/KirinDave/yaws\"},{\"full_name\":\"sr/tasks\",\"html_url\":\"https://github.com/sr/tasks\",\"id\":193,\"name\":\"tasks\",\"url\":\"https://api.github.com/repos/sr/tasks\"},{\"full_name\":\"mattetti/ruby-on-rails-tmbundle\",\"html_url\":\"https://github.com/mattetti/ruby-on-rails-tmbundle\",\"id\":195,\"name\":\"ruby-on-rails-tmbundle\",\"url\":\"https://api.github.com/repos/mattetti/ruby-on-rails-tmbundle\"},{\"full_name\":\"grempe/amazon-ec2\",\"html_url\":\"https://github.com/grempe/amazon-ec2\",\"id\":199,\"name\":\"amazon-ec2\",\"url\":\"https://api.github.com/repos/grempe/amazon-ec2\"},{\"full_name\":\"wayneeseguin/merblogger\",\"html_url\":\"https://github.com/wayneeseguin/merblogger\",\"id\":203,\"name\":\"merblogger\",\"url\":\"https://api.github.com/repos/wayneeseguin/merblogger\"},{\"full_name\":\"wayneeseguin/merbtastic\",\"html_url\":\"https://github.com/wayneeseguin/merbtastic\",\"id\":204,\"name\":\"merbtastic\",\"url\":\"https://api.github.com/repos/wayneeseguin/merbtastic\"},{\"full_name\":\"wayneeseguin/alogr\",\"html_url\":\"https://github.com/wayneeseguin/alogr\",\"id\":205,\"name\":\"alogr\",\"url\":\"https://api.github.com/repos/wayneeseguin/alogr\"},{\"full_name\":\"wayneeseguin/autozest\",\"html_url\":\"https://github.com/wayneeseguin/autozest\",\"id\":206,\"name\":\"autozest\",\"url\":\"https://api.github.com/repos/wayneeseguin/autozest\"},{\"full_name\":\"wayneeseguin/rnginx\",\"html_url\":\"https://github.com/wayneeseguin/rnginx\",\"id\":207,\"name\":\"rnginx\",\"url\":\"https://api.github.com/repos/wayneeseguin/rnginx\"},{\"full_name\":\"wayneeseguin/sequel\",\"html_url\":\"https://github.com/wayneeseguin/sequel\",\"id\":208,\"name\":\"sequel\",\"url\":\"https://api.github.com/repos/wayneeseguin/sequel\"},{\"full_name\":\"bmizerany/simply_versioned\",\"html_url\":\"https://github.com/bmizerany/simply_versioned\",\"id\":211,\"name\":\"simply_versioned\",\"url\":\"https://api.github.com/repos/bmizerany/simply_versioned\"},{\"full_name\":\"peterc/switchpipe\",\"html_url\":\"https://github.com/peterc/switchpipe\",\"id\":212,\"name\":\"switchpipe\",\"url\":\"https://api.github.com/repos/peterc/switchpipe\"},{\"full_name\":\"hornbeck/arc\",\"html_url\":\"https://github.com/hornbeck/arc\",\"id\":213,\"name\":\"arc\",\"url\":\"https://api.github.com/repos/hornbeck/arc\"},{\"full_name\":\"up_the_irons/ebay4r\",\"html_url\":\"https://github.com/up_the_irons/ebay4r\",\"id\":217,\"name\":\"ebay4r\",\"url\":\"https://api.github.com/repos/up_the_irons/ebay4r\"},{\"full_name\":\"wycats/merb-plugins\",\"html_url\":\"https://github.com/wycats/merb-plugins\",\"id\":218,\"name\":\"merb-plugins\",\"url\":\"https://api.github.com/repos/wycats/merb-plugins\"},{\"full_name\":\"up_the_irons/ram\",\"html_url\":\"https://github.com/up_the_irons/ram\",\"id\":220,\"name\":\"ram\",\"url\":\"https://api.github.com/repos/up_the_irons/ram\"},{\"full_name\":\"defunkt/ambitious_activeldap\",\"html_url\":\"https://github.com/defunkt/ambitious_activeldap\",\"id\":230,\"name\":\"ambitious_activeldap\",\"url\":\"https://api.github.com/repos/defunkt/ambitious_activeldap\"},{\"full_name\":\"atmos/fitter_happier\",\"html_url\":\"https://github.com/atmos/fitter_happier\",\"id\":232,\"name\":\"fitter_happier\",\"url\":\"https://api.github.com/repos/atmos/fitter_happier\"},{\"full_name\":\"brosner/oebfare\",\"html_url\":\"https://github.com/brosner/oebfare\",\"id\":237,\"name\":\"oebfare\",\"url\":\"https://api.github.com/repos/brosner/oebfare\"},{\"full_name\":\"up_the_irons/credit_card_tools\",\"html_url\":\"https://github.com/up_the_irons/credit_card_tools\",\"id\":245,\"name\":\"credit_card_tools\",\"url\":\"https://api.github.com/repos/up_the_irons/credit_card_tools\"},{\"full_name\":\"jnicklas/rorem\",\"html_url\":\"https://github.com/jnicklas/rorem\",\"id\":248,\"name\":\"rorem\",\"url\":\"https://api.github.com/repos/jnicklas/rorem\"},{\"full_name\":\"cristibalan/braid\",\"html_url\":\"https://github.com/cristibalan/braid\",\"id\":249,\"name\":\"braid\",\"url\":\"https://api.github.com/repos/cristibalan/braid\"},{\"full_name\":\"jnicklas/uploadcolumn\",\"html_url\":\"https://github.com/jnicklas/uploadcolumn\",\"id\":251,\"name\":\"uploadcolumn\",\"url\":\"https://api.github.com/repos/jnicklas/uploadcolumn\"},{\"full_name\":\"simonjefford/ruby-on-rails-tmbundle\",\"html_url\":\"https://github.com/simonjefford/ruby-on-rails-tmbundle\",\"id\":252,\"name\":\"ruby-on-rails-tmbundle\",\"url\":\"https://api.github.com/repos/simonjefford/ruby-on-rails-tmbundle\"},{\"full_name\":\"chneukirchen/rack-mirror\",\"html_url\":\"https://github.com/chneukirchen/rack-mirror\",\"id\":256,\"name\":\"rack-mirror\",\"url\":\"https://api.github.com/repos/chneukirchen/rack-mirror\"},{\"full_name\":\"chneukirchen/coset-mirror\",\"html_url\":\"https://github.com/chneukirchen/coset-mirror\",\"id\":257,\"name\":\"coset-mirror\",\"url\":\"https://api.github.com/repos/chneukirchen/coset-mirror\"},{\"full_name\":\"drnic/javascript-unittest-tmbundle\",\"html_url\":\"https://github.com/drnic/javascript-unittest-tmbundle\",\"id\":267,\"name\":\"javascript-unittest-tmbundle\",\"url\":\"https://api.github.com/repos/drnic/javascript-unittest-tmbundle\"},{\"full_name\":\"engineyard/eycap\",\"html_url\":\"https://github.com/engineyard/eycap\",\"id\":273,\"name\":\"eycap\",\"url\":\"https://api.github.com/repos/engineyard/eycap\"},{\"full_name\":\"chneukirchen/gitsum\",\"html_url\":\"https://github.com/chneukirchen/gitsum\",\"id\":279,\"name\":\"gitsum\",\"url\":\"https://api.github.com/repos/chneukirchen/gitsum\"},{\"full_name\":\"wayneeseguin/sequel-model\",\"html_url\":\"https://github.com/wayneeseguin/sequel-model\",\"id\":293,\"name\":\"sequel-model\",\"url\":\"https://api.github.com/repos/wayneeseguin/sequel-model\"},{\"full_name\":\"kevinclark/god\",\"html_url\":\"https://github.com/kevinclark/god\",\"id\":305,\"name\":\"god\",\"url\":\"https://api.github.com/repos/kevinclark/god\"},{\"full_name\":\"hornbeck/blerb-core\",\"html_url\":\"https://github.com/hornbeck/blerb-core\",\"id\":307,\"name\":\"blerb-core\",\"url\":\"https://api.github.com/repos/hornbeck/blerb-core\"},{\"full_name\":\"brosner/django-mptt\",\"html_url\":\"https://github.com/brosner/django-mptt\",\"id\":312,\"name\":\"django-mptt\",\"url\":\"https://api.github.com/repos/brosner/django-mptt\"},{\"full_name\":\"technomancy/bus-scheme\",\"html_url\":\"https://github.com/technomancy/bus-scheme\",\"id\":314,\"name\":\"bus-scheme\",\"url\":\"https://api.github.com/repos/technomancy/bus-scheme\"},{\"full_name\":\"Caged/javascript-bits\",\"html_url\":\"https://github.com/Caged/javascript-bits\",\"id\":319,\"name\":\"javascript-bits\",\"url\":\"https://api.github.com/repos/Caged/javascript-bits\"},{\"full_name\":\"Caged/groomlake\",\"html_url\":\"https://github.com/Caged/groomlake\",\"id\":320,\"name\":\"groomlake\",\"url\":\"https://api.github.com/repos/Caged/groomlake\"},{\"full_name\":\"sevenwire/forgery\",\"html_url\":\"https://github.com/sevenwire/forgery\",\"id\":322,\"name\":\"forgery\",\"url\":\"https://api.github.com/repos/sevenwire/forgery\"},{\"full_name\":\"technicalpickles/ambitious-sphinx\",\"html_url\":\"https://github.com/technicalpickles/ambitious-sphinx\",\"id\":324,\"name\":\"ambitious-sphinx\",\"url\":\"https://api.github.com/repos/technicalpickles/ambitious-sphinx\"},{\"full_name\":\"lazyatom/soup\",\"html_url\":\"https://github.com/lazyatom/soup\",\"id\":329,\"name\":\"soup\",\"url\":\"https://api.github.com/repos/lazyatom/soup\"},{\"full_name\":\"josh/rails\",\"html_url\":\"https://github.com/josh/rails\",\"id\":332,\"name\":\"rails\",\"url\":\"https://api.github.com/repos/josh/rails\"},{\"full_name\":\"cdcarter/backpacking\",\"html_url\":\"https://github.com/cdcarter/backpacking\",\"id\":334,\"name\":\"backpacking\",\"url\":\"https://api.github.com/repos/cdcarter/backpacking\"},{\"full_name\":\"jnewland/capsize\",\"html_url\":\"https://github.com/jnewland/capsize\",\"id\":339,\"name\":\"capsize\",\"url\":\"https://api.github.com/repos/jnewland/capsize\"},{\"full_name\":\"bs/starling\",\"html_url\":\"https://github.com/bs/starling\",\"id\":351,\"name\":\"starling\",\"url\":\"https://api.github.com/repos/bs/starling\"},{\"full_name\":\"sr/ape\",\"html_url\":\"https://github.com/sr/ape\",\"id\":360,\"name\":\"ape\",\"url\":\"https://api.github.com/repos/sr/ape\"},{\"full_name\":\"collectiveidea/awesomeness\",\"html_url\":\"https://github.com/collectiveidea/awesomeness\",\"id\":362,\"name\":\"awesomeness\",\"url\":\"https://api.github.com/repos/collectiveidea/awesomeness\"},{\"full_name\":\"collectiveidea/audited\",\"html_url\":\"https://github.com/collectiveidea/audited\",\"id\":363,\"name\":\"audited\",\"url\":\"https://api.github.com/repos/collectiveidea/audited\"},{\"full_name\":\"collectiveidea/acts_as_geocodable\",\"html_url\":\"https://github.com/collectiveidea/acts_as_geocodable\",\"id\":364,\"name\":\"acts_as_geocodable\",\"url\":\"https://api.github.com/repos/collectiveidea/acts_as_geocodable\"},{\"full_name\":\"collectiveidea/acts_as_money\",\"html_url\":\"https://github.com/collectiveidea/acts_as_money\",\"id\":365,\"name\":\"acts_as_money\",\"url\":\"https://api.github.com/repos/collectiveidea/acts_as_money\"},{\"full_name\":\"collectiveidea/calendar_builder\",\"html_url\":\"https://github.com/collectiveidea/calendar_builder\",\"id\":367,\"name\":\"calendar_builder\",\"url\":\"https://api.github.com/repos/collectiveidea/calendar_builder\"},{\"full_name\":\"collectiveidea/clear_empty_attributes\",\"html_url\":\"https://github.com/collectiveidea/clear_empty_attributes\",\"id\":368,\"name\":\"clear_empty_attributes\",\"url\":\"https://api.github.com/repos/collectiveidea/clear_empty_attributes\"},{\"full_name\":\"collectiveidea/css_naked_day\",\"html_url\":\"https://github.com/collectiveidea/css_naked_day\",\"id\":369,\"name\":\"css_naked_day\",\"url\":\"https://api.github.com/repos/collectiveidea/css_naked_day\"}]\n",
//			jsonString:         ``,
//		},
//	}
//	testData := &DataURLs{
//		DataLocation: "https://api.github.com/repositories",
//		URLToPing:    "https://api.github.com",
//	}
//	for _, tc := range tests {
//		testData.URLToPing = tc.url
//		//fmt.Println("Current PATH: ", tc.url)
//		testData.DataLocation = tc.url
//		jsonStr := []byte(`{FullName:"test", HTMLURL:"fakeurl",ID:1231421, Login:"testlog", Name:"testname", URL:"url"}`)
//
//		//fmt.Println("jsonSTRING", tc.jsonString)
//		req, err := http.NewRequest("GET", tc.url, bytes.NewBuffer(jsonStr))
//		// fmt.Println("REQ", req)
//		assert.NoError(err)
//		r := JSONGet(testData)
//		w := httptest.NewRecorder()
//		params := operations.JSONGetParams{HTTPRequest: req, Jsonrepo: []string{tc.url, tc.url}}
//		//fmt.Println("PARAMS:", params.Jsonrepo[0])
//		r(params).WriteResponse(w, runtime.JSONProducer())
//		assert.Equal(tc.expectedStatusCode, w.Code, tc.description)
//		//fmt.Println("tc.expectedStatusCode: ", tc.expectedStatusCode, "w.Code: ", w.Code, "tc.description: ", tc.description)
//		assert.Equal(tc.expectedBody, w.Body.String(), tc.description)
//		//fmt.Println("tc.expectedBody ", tc.expectedBody, "w.Body.String(): ", w.Body.String(), "tc.description", tc.description)
//
//	}
//}
var dp = &models.Jsonrepo{{
	FullName: "artsem",
	HTMLURL:  "artsemURL",
	ID:       12345,
	Login:    "art",
	Name:     "artsemH",
	URL:      "artsem.com",
}}

var tcs = []struct {
	description        string
	reposClient        *models.Jsonrepo
	statusCode         int
	expectedStatusCode int
	expectedBody       string
}{
	{
		description:        "Internal Server Error 500 ",
		reposClient:        &models.Jsonrepo{},
		statusCode:         http.StatusInternalServerError,
		expectedStatusCode: http.StatusInternalServerError,
		expectedBody:       "{\"code\":500,\"message\":\"INTERNAL_SERVER_ERROR\"}\n",
	}, {
		description:        "Resource Not Found",
		reposClient:        &models.Jsonrepo{},
		statusCode:         http.StatusNotFound,
		expectedStatusCode: http.StatusNotFound,
		expectedBody:       "{\"code\":404,\"message\":\"RESOURCE_NOT_FOUND\"}\n",
	}, {
		description:        "No Data Available",
		reposClient:        &models.Jsonrepo{},
		statusCode:         200,
		expectedStatusCode: http.StatusNotFound,
		expectedBody:       "{\"code\":404,\"message\":\"NO_DATA_AVAILABLE\"}\n",
	}, {
		description:        "Get Valid Json",
		reposClient:        dp,
		statusCode:         http.StatusOK,
		expectedStatusCode: http.StatusOK,
		expectedBody:       "[{\"full_name\":\"artsem\",\"html_url\":\"artsemURL\",\"id\":12345,\"login\":\"art\",\"name\":\"artsemH\",\"url\":\"artsem.com\"}]\n",
	},
}

func TestJSONGetMock2(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range tcs {
		resp, _ := json.Marshal(tc.reposClient)
		mockts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(tc.statusCode)
			w.Write(resp)
		}))
		defer mockts.Close()
		testData := &DataURLs{
			DataLocation: mockts.URL,
			URLToPing:    mockts.URL,
		}
		req, err := http.NewRequest("GET", mockts.URL, bytes.NewBuffer(resp))
		assert.NoError(err)
		r := JSONGet(testData)
		w := httptest.NewRecorder()
		params := operations.JSONGetParams{req, []string{mockts.URL, mockts.URL}}
		r(params).WriteResponse(w, runtime.JSONProducer())
		assert.Equal(tc.expectedStatusCode, w.Code, tc.description)
		assert.Equal(tc.expectedBody, w.Body.String(), tc.description)
	}
}
