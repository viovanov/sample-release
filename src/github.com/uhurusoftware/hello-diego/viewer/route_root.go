package viewer

import (
	"net/http"

	"github.com/pivotal-golang/lager"
)

type rootHandler struct {
	logger lager.Logger
}

func NewRootHandler(logger lager.Logger) http.Handler {
	return &rootHandler{
		logger: logger.Session("hello-diego-viewer"),
	}
}

func (handler *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(INDEX_HTML))
}

const INDEX_HTML = `
<html><head>
<style>
.hoverable
{
    fill: transparent;
    stroke:gray; /* Replace with none if you like */
    stroke-width: 4;
    cursor: pointer;
    opacity: 0.2;
}

.hover
{
    stroke:blue;
    opacity: 0.2;
}

svg
{
    border: 1px solid black;
}

</style>

<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>


</head><body>

<svg id="components" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="1160" height="1282">
    <image x="0" y="0" width="100%" height="100%" xlink:href="https://raw.githubusercontent.com/cloudfoundry-incubator/diego-design-notes/master/clickable-diego-overview/clickable-diego-overview.png"/>
    
    <!-- These elements have to be in a specific order, so their Z index is correct. -->
    
    <polygon id="cf_release" class="hoverable" points="9,11,249,11,249,1102,174,1102,174,1194,249,1194,249,1271,9,1271" onclick="window.open('https://github.com/cloudfoundry/cf-release')"/>
    <rect id="etcd_cell" class="hoverable" x="789" y="271" width="360" height="470" onclick="window.open('https://github.com/coreos/etcd')"/>
    <rect id="converger" class="hoverable" x="799" y="651" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/converger')"/>
    <rect id="delta_force" class="hoverable" x="809" y="691" width="120" height="30" onclick="window.open('https://github.com/cloudfoundry-incubator/delta_force')"/>
    <rect id="garden" class="hoverable" x="989" y="471" width="140" height="30" onclick="window.open('https://github.com/cloudfoundry-incubator/garden')"/>
    
    <rect id="executor" class="hoverable" x="989" y="341" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/executor')"/>
    <rect id="etcd_ccbridge" class="hoverable" x="289" y="431" width="160" height="350" onclick="window.open('https://github.com/coreos/etcd')"/>
    <rect id="nsync" class="hoverable" x="299" y="551" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/nsync')"/>
    <rect id="tps" class="hoverable" x="299" y="651" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/tps')"/>    
    <rect id="stager" class="hoverable" x="299" y="451" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/stager')"/>
    
    <rect id="http" class="hoverable" x="956" y="913" width="36" height="23" onclick=""/>
    <rect id="cf_acceptance_tests" class="hoverable" x="180" y="1108" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry/cf-acceptance-tests')"/>
    <rect id="metron" class="hoverable" x="799" y="281" width="340" height="40" onclick="window.open('https://github.com/cloudfoundry/loggregator/tree/develop/src/metron')" />
    <rect id="diego_release" class="hoverable" x="537" y="23" width="141" height="35" onclick="window.open('https://github.com/cloudfoundry-incubator/diego-release')"/>
    
    <polygon id="warden-linux" class="hoverable" points="989,501,1129,501,1129,540,1100,581,989,581" onclick="window.open('https://github.com/cloudfoundry-incubator/warden-linux')"/>
    
    <rect id="cf_release" class="hoverable" x="71" y="23" width="106" height="35" onclick="window.open('https://github.com/cloudfoundry/cf-release')"/>
    <rect id="etcd" class="hoverable" x="539" y="1131" width="160" height="70" onclick="window.open('https://github.com/coreos/etcd')"/>
    <rect id="loggregator" class="hoverable" x="49" y="91" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry/loggregator')"/>
    <rect id="inigo" class="hoverable" x="899" y="781" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/inigo')"/>
    <rect id="linux_circus" class="hoverable" x="297" y="921" width="140" height="120" onclick="window.open('https://github.com/cloudfoundry-incubator/linux-circus')"/>
    <rect id="nats" class="hoverable" x="956" y="891" width="35" height="23" onclick="window.open('https://github.com/apcera/gnatsd')"/>
    <rect id="auctioneer" class="hoverable" x="799" y="551" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/auctioneer')"/>
    <rect id="rep" class="hoverable" x="799" y="341" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/rep')"/>
    <rect id="file_server" class="hoverable" x="299" y="331" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/file-server')"/>
    <rect id="collector" class="hoverable" x="49" y="211" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry/collector')"/>
    <rect id="runtime_metrics_server" class="hoverable" x="299" y="211" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/runtime-metrics-server')"/>
    <rect id="router" class="hoverable" x="49" y="811" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry/gorouter')"/>
    <rect id="route_emmiter" class="hoverable" x="299" y="811" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry-incubator/route-emitter')"/>
    <rect id="cloud_controller_ng" class="hoverable" x="49" y="511" width="140" height="80" onclick="window.open('https://github.com/cloudfoundry/cloud_controller_ng')"/>
    <rect id="runtime_schema" class="hoverable" x="539" y="171" width="160" height="870" onclick="window.open('https://github.com/cloudfoundry-incubator/runtime-schema')"/>
    <rect id="storeadapter" class="hoverable" x="539" y="1041" width="160" height="40" onclick="window.open('https://github.com/cloudfoundry/storeadapter')"/>
    <polygon id="auction" class="hoverable" points="799,421,799,551,939,551,939,421,799,421" onclick="window.open('https://github.com/cloudfoundry-incubator/auction')"/>    
</svg>


<script type="text/javascript">

function pingComponent(item)
{
    var element = $("#" + item)
    var old_stroke = element.css('stroke');

    element.css("stroke", "#FF0000");
    element.animate({opacity:'1'}, {
        duration: 1000,
        complete: function () {
            element.animate({opacity:'0.2'}, {
                duration: 2000,
                complete: function () {
                    element.css('stroke', 'gray');
                }
            })
        }
    });
}

$("rect, polygon").hover(
    function() {
        $(this).css('stroke', 'blue');
    }, 
    function() {
        $(this).css('stroke', 'gray');
    }
);

setInterval(function () 
{
	$.ajax({
            url: "/update",
            dataType: 'text',
            success: function(data){            
                $("#components").children("rect, polygon").each (function() {
    				var element = this;
					if (data.indexOf(element.id) > -1)
					{
						pingComponent(element.id);						
					}
				});
            }
        });
}, 100);

</script>

</body></html>
`
