	var path2 = [[0,21],[1,21],[2,21],[3,21],[3,22],[3,23],[3,24],[3,25],[3,26],[3,27],[3,28],[3,29],[3,30],[3,31],[3,32],[4,32],[5,32],[6,32],[7,32],[8,32],[9,32],[10,32],[11,32],[12,32],[13,32],[14,32],[15,32],[16,32],[17,32],[18,32],[19,32],[20,32],[20,33],[21,33],[21,34],[21,35],[21,36],[21,37],[21,38],[21,39],[21,40],[21,41],[21,42],[21,43],[21,44],[22,44],[23,44],[24,44],[25,44],[26,44],[26,45],[27,45],[27,46],[27,47],[27,48],[27,49],[27,50],[26,50],[25,50],[24,50],[23,50],[23,51],[23,52],[23,53],[23,54],[23,55],[23,56],[24,56],[25,56],[26,56],[27,56],[27,57],[27,58],[26,58],[26,59],[26,60],[25,60],[24,60],[23,60],[22,60],[21,60],[20,60],[20,61]];

	var path = [[20,61],[15,66],[9,70],[8,72],[3,21],[0,20]];

	var popped = [[0,20],[3,21],[3,23],[3,28],[3,30],[5,21],[5,23],[3,35],[5,30],[5,34],[5,36],[7,24],[3,41],[2,18],[2,36],[8,26],[8,30],[8,30],[3,21],[9,24],[5,44],[3,46],[8,36],[3,48],[9,38],[9,40],[5,48],[3,17],[11,27],[9,44],[0,18],[5,32],[6,26],[11,38],[3,39],[9,47],[11,40],[0,36],[12,30],[11,42],[5,51],[3,52],[9,49],[0,38],[5,17],[2,36],[11,48],[10,50],[3,44],[5,27],[5,27],[2,38],[7,38],[5,30],[5,54],[14,32],[5,32],[3,39],[14,44],[3,55],[5,36],[3,31],[9,38],[15,42],[5,56],[15,48],[3,35],[15,50],[2,38],[3,50],[16,30],[15,52],[59,42],[9,45],[59,35],[3,57],[16,44],[0,20],[6,26],[17,43],[5,58],[16,56],[17,52],[11,46],[17,54],[5,39],[12,44],[3,14],[8,26],[14,53],[18,30],[3,17],[58,44],[14,56],[6,60],[8,60],[5,14],[11,46],[2,18],[5,39],[57,42],[13,54],[20,32],[3,44],[15,50],[11,59],[59,42],[3,4],[12,50],[7,14],[3,1],[12,30],[21,33],[21,36],[11,42],[21,41],[3,62],[7,62],[3,94],[5,10],[57,34],[5,7],[12,44],[11,48],[5,3],[57,36],[22,30],[9,15],[9,17],[9,19],[7,38],[5,92],[2,93],[7,12],[19,36],[6,60],[6,90],[7,24],[23,42],[3,41],[5,64],[5,14],[17,49],[55,42],[11,54],[7,92],[13,54],[11,18],[3,78],[9,56],[3,59],[5,80],[55,44],[9,12],[11,56],[11,56],[24,36],[1,92],[9,8],[9,5],[57,44],[1,90],[5,78],[9,0],[12,20],[56,33],[5,77],[19,30],[9,90],[21,36],[3,75],[21,39],[5,62],[5,62],[17,36],[9,87],[5,12],[5,12],[10,0],[25,30],[5,9],[11,58],[3,90],[3,73],[5,74],[11,10],[11,8],[14,53],[11,3],[5,80],[3,72],[9,15],[11,91],[5,92],[11,88],[26,36],[26,38],[12,2],[26,42],[21,33],[26,44],[9,54],[2,80],[10,54],[23,40],[3,70],[12,86],[15,18],[5,58],[15,20],[15,23],[59,92],[5,70],[5,6],[17,55],[27,30],[27,32],[5,23],[9,56],[5,95],[11,27],[54,33],[15,52],[24,36],[24,38],[1,92],[3,68],[7,74],[9,7],[27,45],[5,68],[9,2],[27,47],[56,33],[19,38],[9,74],[9,68],[17,21],[17,24],[23,40],[23,42],[8,72],[15,2],[28,38],[11,68],[57,34],[15,86],[13,20],[11,74],[11,12],[11,15],[20,32],[18,26],[9,70],[24,32],[3,65],[53,35],[29,36],[24,48],[2,93],[7,14],[6,72],[58,38],[59,27],[19,24],[9,90],[59,94],[53,37],[17,0],[0,78],[23,50],[59,92],[29,46],[23,53],[51,44],[29,48],[5,74],[13,68],[11,18],[20,26],[3,78],[18,2],[22,48],[9,70],[59,75],[17,64],[9,12],[24,32],[24,38],[59,35],[27,45],[15,68],[56,90],[4,86],[6,72],[21,24],[0,36],[19,86],[8,60],[21,50],[21,52],[24,54],[15,66]];

	var expanded = [[3,23],[3,28],[3,30],[3,35],[3,41],[3,46],[3,48],[3,52],[3,55],[3,57],[3,62],[3,68],[3,70],[3,72],[3,75],[3,78],[3,14],[3,4],[3,1],[3,94],[5,21],[5,23],[5,30],[5,34],[5,36],[5,44],[5,48],[5,51],[5,54],[5,56],[5,58],[5,64],[5,68],[5,70],[5,77],[5,80],[5,17],[5,14],[5,10],[5,7],[5,3],[5,92],[3,21],[2,18],[9,24],[8,26],[8,30],[16,30],[18,30],[22,30],[25,30],[27,30],[33,30],[14,32],[20,32],[27,32],[32,32],[35,32],[0,36],[8,36],[0,38],[9,38],[57,42],[55,42],[45,42],[58,44],[55,44],[51,44],[47,44],[8,60],[7,62],[8,72],[7,74],[9,74],[11,74],[0,78],[55,78],[52,78],[49,78],[46,78],[42,78],[35,78],[2,80],[57,80],[54,80],[44,80],[35,80],[5,80],[3,21],[7,24],[3,31],[2,36],[5,27],[8,26],[21,36],[21,41],[23,42],[24,36],[27,27],[27,21],[27,17],[27,14],[27,9],[27,4],[29,27],[29,23],[29,16],[29,10],[29,8],[29,6],[29,4],[33,27],[33,25],[33,22],[33,20],[33,17],[33,11],[33,7],[33,4],[33,2],[33,94],[33,92],[33,87],[33,79],[33,76],[33,71],[33,69],[33,62],[33,52],[33,49],[33,46],[35,32],[35,26],[35,20],[35,16],[35,14],[35,8],[35,6],[35,3],[35,0],[35,92],[35,87],[35,82],[35,80],[35,75],[35,69],[35,65],[35,62],[35,58],[35,55],[35,51],[35,44],[5,32],[9,40],[9,44],[9,47],[9,49],[11,38],[11,40],[11,42],[11,48],[8,30],[11,27],[3,39],[59,42],[0,18],[3,17],[59,27],[57,34],[57,27],[6,26],[5,30],[5,32],[7,24],[0,20],[5,39],[3,44],[5,36],[7,38],[15,42],[14,44],[16,44],[15,48],[10,50],[15,50],[9,38],[7,14],[7,12],[9,12],[7,92],[59,92],[6,90],[9,90],[1,90],[56,90],[53,90],[12,30],[35,30],[5,27],[2,38],[9,45],[2,36],[59,35],[17,43],[3,50],[2,38],[3,35],[11,46],[15,52],[17,52],[17,54],[45,42],[6,26],[3,39],[5,39],[4,90],[12,44],[57,37],[11,42],[11,48],[12,50],[11,54],[14,56],[11,56],[16,56],[12,30],[15,50],[14,53],[54,33],[51,30],[52,26],[47,26],[58,24],[55,24],[6,60],[17,49],[12,44],[11,46],[3,17],[13,54],[2,18],[3,44],[9,54],[11,59],[11,56],[6,60],[9,56],[9,17],[9,19],[9,8],[9,5],[9,0],[11,18],[11,10],[11,8],[11,3],[7,38],[59,42],[11,58],[9,56],[21,33],[3,41],[3,12],[5,14],[9,15],[11,27],[19,36],[17,36],[26,36],[29,36],[26,38],[28,38],[26,42],[31,42],[33,42],[36,42],[39,42],[26,44],[31,44],[37,44],[21,33],[9,36],[21,39],[3,59],[5,62],[2,93],[5,12],[57,36],[56,33],[5,9],[5,6],[19,30],[15,18],[12,20],[15,20],[9,15],[5,95],[9,87],[11,91],[11,88],[1,92],[5,12],[21,36],[5,58],[5,23],[23,40],[27,47],[29,46],[29,48],[33,60],[33,65],[33,78],[33,83],[33,89],[33,0],[33,9],[33,14],[33,24],[33,30],[35,47],[35,54],[35,57],[35,60],[35,71],[35,78],[35,84],[35,10],[35,22],[35,29],[39,39],[41,43],[41,40],[5,62],[13,54],[5,92],[14,53],[11,15],[15,23],[17,21],[17,24],[3,73],[3,65],[59,75],[59,71],[59,64],[59,58],[59,56],[59,54],[59,52],[57,75],[57,69],[57,65],[57,62],[57,57],[57,52],[54,84],[53,90],[51,84],[51,89],[38,74],[10,54],[5,78],[5,74],[57,44],[15,2],[18,2],[23,2],[26,2],[10,0],[17,0],[21,0],[28,0],[17,55],[8,62],[24,32],[59,94],[59,0],[57,95],[53,88],[51,87],[51,81],[9,12],[9,7],[3,90],[9,2],[53,37],[51,31],[51,35],[15,86],[19,86],[23,86],[20,84],[19,38],[9,90],[4,86],[9,68],[11,68],[11,12],[15,52],[11,7],[12,2],[7,14],[12,86],[24,36],[24,38],[27,7],[27,12],[27,16],[27,25],[27,30],[29,15],[29,20],[29,26],[29,29],[23,42],[20,32],[27,45],[24,38],[21,79],[23,83],[11,18],[13,20],[19,24],[21,24],[18,26],[20,26],[15,20],[1,92],[52,0],[57,2],[53,2],[6,72],[24,2],[21,2],[12,2],[18,0],[13,0],[24,32],[8,60],[10,24],[56,33],[53,35],[23,40],[2,93],[5,74],[22,48],[23,50],[27,45],[24,48],[57,34],[9,70],[21,20],[21,18],[23,26],[23,23],[23,21],[23,17],[23,15],[23,12],[6,72],[9,70],[11,72],[15,61],[17,64],[17,62],[21,71],[24,71],[58,38],[12,86],[13,68],[15,68],[20,68],[15,66],[44,36],[50,38],[46,38],[42,38],[23,53],[21,50],[21,52],[3,94],[59,35],[17,24],[59,92],[53,35],[3,78],[59,77],[26,50],[28,54],[21,56],[27,56],[29,44],[23,50],[24,54],[11,68],[24,48],[8,72],[59,77],[17,66],[29,37],[0,36],[26,44],[5,70],[15,14],[18,12],[16,86],[27,58],[29,60],[11,66],[19,62],[21,62],[27,62],[29,62],[20,61]]

	var water = [[0,1],[0,2],[0,3],[0,4],[0,5],[0,10],[0,11],[0,13],[0,14],[0,15],[0,16],[0,22],[0,23],[0,24],[0,28],[0,29],[0,30],[0,31],[0,32],[0,33],[0,34],[0,40],[0,45],[0,46],[0,47],[0,48],[0,53],[0,55],[0,57],[0,59],[0,60],[0,61],[0,62],[0,65],[0,66],[0,67],[0,72],[0,73],[0,76],[0,81],[0,82],[0,83],[0,86],[0,87],[0,88],[0,89],[0,93],[0,95],[1,0],[1,1],[1,2],[1,3],[1,4],[1,5],[1,6],[1,7],[1,8],[1,9],[1,10],[1,11],[1,12],[1,13],[1,14],[1,15],[1,16],[1,17],[1,22],[1,23],[1,24],[1,25],[1,26],[1,27],[1,28],[1,29],[1,30],[1,31],[1,32],[1,33],[1,34],[1,35],[1,39],[1,40],[1,45],[1,46],[1,47],[1,48],[1,49],[1,50],[1,51],[1,52],[1,53],[1,54],[1,55],[1,56],[1,57],[1,58],[1,59],[1,60],[1,61],[1,62],[1,63],[1,64],[1,65],[1,66],[1,67],[1,68],[1,69],[1,70],[1,71],[1,72],[1,73],[1,74],[1,75],[1,76],[1,77],[1,81],[1,82],[1,83],[1,84],[1,85],[1,86],[1,87],[1,88],[1,94],[1,95],[2,0],[2,2],[2,3],[2,5],[2,6],[2,7],[2,8],[2,9],[2,10],[2,11],[2,15],[2,16],[2,22],[2,24],[2,25],[2,26],[2,27],[2,29],[2,32],[2,33],[2,34],[2,40],[2,45],[2,47],[2,51],[2,53],[2,54],[2,56],[2,60],[2,61],[2,66],[2,67],[2,69],[2,71],[2,74],[2,76],[2,77],[2,82],[2,83],[2,84],[2,85],[2,86],[2,87],[2,88],[2,89],[2,95],[3,81],[3,82],[3,83],[3,87],[3,88],[4,81],[4,82],[4,83],[4,88],[5,82],[5,83],[5,88],[5,89],[6,4],[6,5],[6,8],[6,11],[6,15],[6,16],[6,18],[6,19],[6,20],[6,22],[6,28],[6,29],[6,33],[6,35],[6,40],[6,41],[6,42],[6,43],[6,45],[6,46],[6,47],[6,49],[6,50],[6,52],[6,53],[6,55],[6,57],[6,63],[6,65],[6,66],[6,67],[6,69],[6,75],[6,76],[6,79],[6,81],[6,82],[6,83],[6,88],[6,93],[6,94],[7,0],[7,1],[7,2],[7,3],[7,4],[7,5],[7,6],[7,7],[7,8],[7,9],[7,10],[7,16],[7,17],[7,18],[7,19],[7,20],[7,21],[7,22],[7,27],[7,28],[7,29],[7,33],[7,34],[7,35],[7,40],[7,41],[7,42],[7,43],[7,44],[7,45],[7,46],[7,47],[7,48],[7,49],[7,50],[7,51],[7,52],[7,53],[7,54],[7,55],[7,56],[7,57],[7,58],[7,59],[7,64],[7,65],[7,66],[7,67],[7,68],[7,69],[7,70],[7,71],[7,76],[7,77],[7,78],[7,79],[7,80],[7,81],[7,82],[7,83],[7,88],[7,89],[7,94],[7,95],[8,1],[8,6],[8,9],[8,10],[8,11],[8,16],[8,18],[8,21],[8,22],[8,23],[8,33],[8,34],[8,39],[8,41],[8,42],[8,43],[8,46],[8,48],[8,51],[8,52],[8,53],[8,55],[8,64],[8,65],[8,66],[8,69],[8,75],[8,76],[8,77],[8,78],[8,79],[8,80],[8,81],[8,82],[8,83],[8,88],[8,89],[8,94],[8,95],[9,21],[9,22],[9,33],[9,34],[9,51],[9,52],[9,53],[9,63],[9,64],[9,76],[9,77],[9,82],[9,94],[9,95],[10,21],[10,22],[10,33],[10,34],[10,35],[10,52],[10,63],[10,64],[10,75],[10,76],[10,77],[10,81],[10,82],[10,94],[11,21],[11,22],[11,23],[11,33],[11,34],[11,35],[11,52],[11,63],[11,64],[11,76],[11,77],[11,82],[11,93],[11,94],[12,4],[12,5],[12,6],[12,9],[12,11],[12,16],[12,17],[12,22],[12,23],[12,24],[12,25],[12,26],[12,33],[12,34],[12,35],[12,36],[12,37],[12,39],[12,41],[12,47],[12,52],[12,53],[12,57],[12,60],[12,61],[12,62],[12,63],[12,64],[12,65],[12,69],[12,70],[12,71],[12,75],[12,76],[12,77],[12,82],[12,89],[12,90],[12,92],[12,93],[12,94],[13,3],[13,4],[13,5],[13,6],[13,7],[13,8],[13,9],[13,10],[13,11],[13,12],[13,13],[13,14],[13,15],[13,16],[13,17],[13,22],[13,23],[13,24],[13,25],[13,26],[13,27],[13,28],[13,29],[13,33],[13,34],[13,35],[13,36],[13,37],[13,38],[13,39],[13,40],[13,41],[13,45],[13,46],[13,47],[13,51],[13,52],[13,57],[13,58],[13,59],[13,60],[13,61],[13,62],[13,63],[13,64],[13,65],[13,70],[13,71],[13,72],[13,73],[13,74],[13,75],[13,76],[13,81],[13,82],[13,87],[13,88],[13,89],[13,90],[13,91],[13,92],[13,93],[13,94],[14,3],[14,4],[14,5],[14,6],[14,8],[14,9],[14,11],[14,12],[14,14],[14,15],[14,16],[14,17],[14,21],[14,22],[14,25],[14,26],[14,27],[14,28],[14,29],[14,34],[14,35],[14,36],[14,37],[14,38],[14,39],[14,40],[14,41],[14,46],[14,47],[14,51],[14,58],[14,59],[14,62],[14,63],[14,64],[14,65],[14,69],[14,70],[14,71],[14,72],[14,73],[14,74],[14,77],[14,81],[14,82],[14,87],[14,88],[14,89],[14,91],[14,92],[14,93],[14,94],[14,95],[15,4],[15,5],[15,16],[15,27],[15,28],[15,29],[15,33],[15,34],[15,35],[15,39],[15,40],[15,45],[15,46],[15,57],[15,58],[15,70],[15,71],[15,82],[15,83],[15,88],[15,94],[15,95],[16,3],[16,4],[16,5],[16,15],[16,16],[16,17],[16,27],[16,28],[16,33],[16,34],[16,39],[16,40],[16,46],[16,58],[16,59],[16,69],[16,70],[16,71],[16,81],[16,82],[16,83],[16,88],[16,89],[16,94],[16,95],[17,3],[17,4],[17,15],[17,16],[17,17],[17,27],[17,28],[17,29],[17,33],[17,34],[17,39],[17,40],[17,45],[17,46],[17,47],[17,58],[17,59],[17,69],[17,70],[17,71],[17,81],[17,82],[17,83],[17,87],[17,88],[17,89],[17,94],[18,4],[18,9],[18,15],[18,16],[18,17],[18,18],[18,19],[18,20],[18,22],[18,23],[18,28],[18,33],[18,34],[18,35],[18,39],[18,40],[18,41],[18,42],[18,45],[18,46],[18,47],[18,50],[18,51],[18,53],[18,56],[18,57],[18,58],[18,59],[18,63],[18,65],[18,69],[18,70],[18,75],[18,81],[18,82],[18,83],[18,87],[18,88],[18,89],[18,93],[18,94],[19,4],[19,5],[19,10],[19,11],[19,15],[19,16],[19,17],[19,18],[19,19],[19,20],[19,21],[19,22],[19,27],[19,28],[19,33],[19,34],[19,40],[19,41],[19,42],[19,43],[19,44],[19,45],[19,46],[19,47],[19,48],[19,49],[19,50],[19,51],[19,52],[19,53],[19,54],[19,55],[19,56],[19,57],[19,58],[19,64],[19,65],[19,69],[19,70],[19,75],[19,76],[19,77],[19,78],[19,79],[19,80],[19,81],[19,82],[19,83],[19,88],[19,89],[19,93],[19,94],[19,95],[20,4],[20,5],[20,9],[20,10],[20,11],[20,19],[20,21],[20,22],[20,23],[20,28],[20,29],[20,34],[20,35],[20,40],[20,42],[20,43],[20,44],[20,45],[20,46],[20,47],[20,48],[20,49],[20,51],[20,56],[20,57],[20,58],[20,59],[20,63],[20,64],[20,65],[20,70],[20,75],[20,76],[20,77],[20,78],[20,80],[20,81],[20,82],[20,88],[20,89],[20,94],[20,95],[21,4],[21,9],[21,10],[21,28],[21,29],[21,45],[21,46],[21,47],[21,58],[21,64],[21,65],[21,75],[21,76],[21,87],[21,88],[21,89],[21,94],[22,3],[22,4],[22,5],[22,10],[22,11],[22,28],[22,45],[22,46],[22,57],[22,58],[22,64],[22,65],[22,76],[22,77],[22,87],[22,88],[22,89],[22,93],[22,94],[22,95],[23,4],[23,5],[23,10],[23,11],[23,28],[23,29],[23,45],[23,46],[23,47],[23,58],[23,64],[23,65],[23,75],[23,76],[23,88],[23,93],[23,94],[23,95],[24,4],[24,5],[24,6],[24,7],[24,9],[24,10],[24,11],[24,13],[24,14],[24,16],[24,18],[24,22],[24,24],[24,25],[24,27],[24,28],[24,29],[24,41],[24,45],[24,46],[24,51],[24,52],[24,57],[24,58],[24,64],[24,65],[24,66],[24,76],[24,84],[24,87],[24,88],[24,93],[24,94],[24,95],[25,3],[25,4],[25,5],[25,6],[25,7],[25,8],[25,9],[25,10],[25,11],[25,12],[25,13],[25,14],[25,15],[25,16],[25,17],[25,18],[25,19],[25,20],[25,21],[25,22],[25,23],[25,24],[25,25],[25,26],[25,27],[25,28],[25,33],[25,34],[25,35],[25,39],[25,40],[25,41],[25,45],[25,46],[25,51],[25,52],[25,53],[25,57],[25,58],[25,59],[25,64],[25,65],[25,66],[25,67],[25,68],[25,69],[25,70],[25,76],[25,82],[25,83],[25,84],[25,85],[25,86],[25,87],[25,88],[25,94],[25,95],[26,5],[26,6],[26,10],[26,11],[26,15],[26,18],[26,19],[26,20],[26,22],[26,23],[26,24],[26,28],[26,29],[26,33],[26,34],[26,40],[26,46],[26,52],[26,53],[26,57],[26,63],[26,64],[26,65],[26,66],[26,67],[26,68],[26,69],[26,70],[26,71],[26,76],[26,83],[26,86],[26,87],[26,88],[26,89],[26,93],[26,94],[26,95],[27,34],[27,35],[27,39],[27,40],[27,41],[27,52],[27,53],[27,64],[27,65],[27,69],[27,70],[27,87],[27,88],[27,94],[27,95],[28,33],[28,34],[28,35],[28,40],[28,41],[28,51],[28,52],[28,63],[28,64],[28,69],[28,70],[28,88],[28,89],[28,94],[29,33],[29,34],[29,39],[29,40],[29,41],[29,51],[29,52],[29,64],[29,69],[29,70],[29,71],[29,88],[29,93],[29,94],[30,0],[30,5],[30,7],[30,9],[30,11],[30,12],[30,13],[30,14],[30,17],[30,18],[30,19],[30,24],[30,25],[30,28],[30,33],[30,34],[30,35],[30,38],[30,39],[30,40],[30,41],[30,45],[30,47],[30,49],[30,50],[30,51],[30,52],[30,53],[30,58],[30,59],[30,61],[30,62],[30,63],[30,64],[30,70],[30,71],[30,72],[30,76],[30,77],[30,78],[30,79],[30,80],[30,81],[30,82],[30,88],[30,93],[30,94],[30,95],[31,0],[31,1],[31,2],[31,3],[31,4],[31,5],[31,6],[31,7],[31,8],[31,9],[31,10],[31,11],[31,12],[31,13],[31,14],[31,15],[31,16],[31,17],[31,18],[31,19],[31,20],[31,21],[31,22],[31,23],[31,24],[31,25],[31,26],[31,27],[31,28],[31,29],[31,33],[31,34],[31,35],[31,36],[31,37],[31,38],[31,39],[31,40],[31,46],[31,47],[31,48],[31,49],[31,50],[31,51],[31,52],[31,53],[31,54],[31,55],[31,56],[31,57],[31,58],[31,59],[31,60],[31,61],[31,62],[31,63],[31,64],[31,65],[31,70],[31,71],[31,72],[31,73],[31,74],[31,75],[31,76],[31,77],[31,78],[31,79],[31,80],[31,81],[31,82],[31,83],[31,87],[31,88],[31,93],[31,94],[31,95],[32,3],[32,5],[32,6],[32,8],[32,12],[32,13],[32,18],[32,19],[32,21],[32,23],[32,26],[32,28],[32,29],[32,34],[32,35],[32,36],[32,37],[32,38],[32,39],[32,40],[32,41],[32,47],[32,48],[32,50],[32,51],[32,53],[32,54],[32,55],[32,56],[32,57],[32,58],[32,59],[32,63],[32,64],[32,70],[32,72],[32,73],[32,74],[32,75],[32,77],[32,80],[32,81],[32,82],[32,88],[32,93],[32,95],[33,33],[33,34],[33,35],[33,39],[33,40],[34,33],[34,34],[34,35],[34,40],[35,34],[35,35],[35,40],[35,41],[36,1],[36,2],[36,4],[36,5],[36,7],[36,9],[36,15],[36,17],[36,18],[36,19],[36,21],[36,27],[36,28],[36,31],[36,33],[36,34],[36,35],[36,40],[36,45],[36,46],[36,52],[36,53],[36,56],[36,59],[36,63],[36,64],[36,66],[36,67],[36,68],[36,70],[36,76],[36,77],[36,81],[36,83],[36,88],[36,89],[36,90],[36,91],[36,93],[36,94],[36,95],[37,0],[37,1],[37,2],[37,3],[37,4],[37,5],[37,6],[37,7],[37,8],[37,9],[37,10],[37,11],[37,16],[37,17],[37,18],[37,19],[37,20],[37,21],[37,22],[37,23],[37,28],[37,29],[37,30],[37,31],[37,32],[37,33],[37,34],[37,35],[37,40],[37,41],[37,46],[37,47],[37,48],[37,49],[37,50],[37,51],[37,52],[37,53],[37,54],[37,55],[37,56],[37,57],[37,58],[37,64],[37,65],[37,66],[37,67],[37,68],[37,69],[37,70],[37,75],[37,76],[37,77],[37,81],[37,82],[37,83],[37,88],[37,89],[37,90],[37,91],[37,92],[37,93],[37,94],[37,95],[38,0],[38,3],[38,4],[38,5],[38,7],[38,16],[38,17],[38,18],[38,21],[38,27],[38,28],[38,29],[38,30],[38,31],[38,32],[38,33],[38,34],[38,35],[38,40],[38,41],[38,46],[38,47],[38,49],[38,54],[38,57],[38,58],[38,59],[38,64],[38,66],[38,69],[38,70],[38,71],[38,81],[38,82],[38,87],[38,89],[38,90],[38,91],[38,94],[39,3],[39,4],[39,5],[39,15],[39,16],[39,28],[39,29],[39,34],[39,46],[39,47],[39,69],[39,70],[39,81],[39,82],[40,4],[40,15],[40,16],[40,27],[40,28],[40,29],[40,33],[40,34],[40,46],[40,69],[40,70],[40,81],[40,82],[40,83],[41,4],[41,15],[41,16],[41,28],[41,29],[41,34],[41,45],[41,46],[41,69],[41,70],[41,71],[41,81],[41,82],[41,83],[42,4],[42,5],[42,9],[42,12],[42,13],[42,14],[42,15],[42,16],[42,17],[42,21],[42,22],[42,23],[42,27],[42,28],[42,29],[42,34],[42,41],[42,42],[42,44],[42,45],[42,46],[42,52],[42,53],[42,54],[42,57],[42,59],[42,64],[42,65],[42,70],[42,71],[42,72],[42,73],[42,74],[42,81],[42,82],[42,83],[42,84],[42,85],[42,87],[42,89],[42,95],[43,3],[43,4],[43,9],[43,10],[43,11],[43,12],[43,13],[43,14],[43,15],[43,16],[43,17],[43,22],[43,23],[43,24],[43,25],[43,26],[43,27],[43,28],[43,33],[43,34],[43,39],[43,40],[43,41],[43,42],[43,43],[43,44],[43,45],[43,46],[43,51],[43,52],[43,53],[43,54],[43,55],[43,56],[43,57],[43,58],[43,59],[43,60],[43,61],[43,62],[43,63],[43,64],[43,65],[43,70],[43,71],[43,72],[43,73],[43,74],[43,75],[43,76],[43,77],[43,81],[43,82],[43,83],[43,84],[43,85],[43,86],[43,87],[43,88],[43,89],[43,93],[43,94],[43,95],[44,3],[44,10],[44,11],[44,14],[44,15],[44,16],[44,17],[44,21],[44,22],[44,23],[44,24],[44,25],[44,26],[44,29],[44,33],[44,34],[44,39],[44,40],[44,41],[44,43],[44,44],[44,45],[44,46],[44,47],[44,51],[44,52],[44,53],[44,54],[44,56],[44,57],[44,59],[44,60],[44,62],[44,63],[44,64],[44,65],[44,69],[44,70],[44,73],[44,74],[44,75],[44,76],[44,77],[44,82],[44,83],[44,84],[44,85],[44,86],[44,87],[44,88],[44,89],[44,94],[44,95],[45,9],[45,10],[45,22],[45,23],[45,34],[45,35],[45,40],[45,46],[45,47],[45,52],[45,53],[45,64],[45,75],[45,76],[45,77],[45,81],[45,82],[45,83],[45,87],[45,88],[45,93],[45,94],[46,10],[46,11],[46,21],[46,22],[46,23],[46,33],[46,34],[46,35],[46,40],[46,41],[46,46],[46,47],[46,51],[46,52],[46,53],[46,63],[46,64],[46,65],[46,75],[46,76],[46,81],[46,82],[46,87],[46,88],[46,94],[47,10],[47,11],[47,21],[47,22],[47,23],[47,33],[47,34],[47,35],[47,39],[47,40],[47,41],[47,46],[47,51],[47,52],[47,63],[47,64],[47,65],[47,75],[47,76],[47,77],[47,81],[47,82],[47,87],[47,88],[47,93],[47,94],[47,95],[48,2],[48,3],[48,5],[48,8],[48,9],[48,10],[48,11],[48,15],[48,17],[48,21],[48,22],[48,27],[48,33],[48,34],[48,35],[48,39],[48,40],[48,41],[48,45],[48,46],[48,52],[48,57],[48,63],[48,64],[48,65],[48,66],[48,67],[48,68],[48,70],[48,71],[48,76],[48,81],[48,82],[48,83],[48,87],[48,88],[48,89],[48,90],[48,93],[48,94],[48,95],[49,0],[49,1],[49,2],[49,3],[49,4],[49,5],[49,6],[49,7],[49,8],[49,9],[49,10],[49,16],[49,17],[49,21],[49,22],[49,27],[49,28],[49,29],[49,30],[49,31],[49,32],[49,33],[49,34],[49,35],[49,40],[49,41],[49,45],[49,46],[49,47],[49,52],[49,53],[49,58],[49,59],[49,63],[49,64],[49,65],[49,66],[49,67],[49,68],[49,69],[49,70],[49,75],[49,76],[49,81],[49,82],[49,88],[49,89],[49,90],[49,91],[49,92],[49,93],[49,94],[49,95],[50,0],[50,1],[50,3],[50,8],[50,9],[50,10],[50,11],[50,15],[50,16],[50,17],[50,22],[50,27],[50,28],[50,29],[50,30],[50,32],[50,33],[50,34],[50,40],[50,41],[50,46],[50,47],[50,52],[50,53],[50,57],[50,58],[50,59],[50,67],[50,69],[50,70],[50,71],[50,76],[50,77],[50,82],[50,83],[50,88],[50,90],[50,91],[50,92],[50,93],[50,94],[50,95],[51,10],[51,16],[51,17],[51,27],[51,28],[51,39],[51,40],[51,41],[51,46],[51,52],[51,57],[51,58],[51,76],[51,77],[51,93],[51,94],[51,95],[52,9],[52,10],[52,16],[52,17],[52,28],[52,29],[52,39],[52,40],[52,41],[52,45],[52,46],[52,47],[52,51],[52,52],[52,53],[52,58],[52,59],[52,76],[52,93],[52,94],[53,10],[53,16],[53,17],[53,27],[53,28],[53,40],[53,45],[53,46],[53,47],[53,52],[53,53],[53,58],[53,59],[53,76],[53,77],[53,93],[53,94],[53,95],[54,3],[54,4],[54,9],[54,10],[54,16],[54,17],[54,18],[54,28],[54,36],[54,39],[54,40],[54,45],[54,46],[54,47],[54,52],[54,53],[54,54],[54,55],[54,57],[54,58],[54,59],[54,61],[54,62],[54,64],[54,66],[54,70],[54,72],[54,73],[54,75],[54,76],[54,77],[54,89],[54,93],[54,94],[55,3],[55,4],[55,5],[55,9],[55,10],[55,11],[55,16],[55,17],[55,18],[55,19],[55,20],[55,21],[55,22],[55,28],[55,34],[55,35],[55,36],[55,37],[55,38],[55,39],[55,40],[55,46],[55,47],[55,51],[55,52],[55,53],[55,54],[55,55],[55,56],[55,57],[55,58],[55,59],[55,60],[55,61],[55,62],[55,63],[55,64],[55,65],[55,66],[55,67],[55,68],[55,69],[55,70],[55,71],[55,72],[55,73],[55,74],[55,75],[55,76],[55,81],[55,82],[55,83],[55,87],[55,88],[55,89],[55,93],[55,94],[56,4],[56,5],[56,9],[56,15],[56,16],[56,17],[56,18],[56,19],[56,20],[56,21],[56,22],[56,23],[56,28],[56,35],[56,38],[56,39],[56,40],[56,41],[56,45],[56,46],[56,47],[56,53],[56,54],[56,58],[56,59],[56,63],[56,66],[56,67],[56,68],[56,70],[56,71],[56,72],[56,76],[56,77],[56,81],[56,82],[56,88],[56,94],[57,4],[57,5],[57,16],[57,17],[57,21],[57,22],[57,39],[57,40],[57,46],[57,47],[57,82],[57,83],[57,87],[57,88],[57,89],[58,3],[58,4],[58,15],[58,16],[58,21],[58,22],[58,40],[58,41],[58,46],[58,81],[58,82],[58,83],[58,88],[58,89],[59,3],[59,4],[59,16],[59,21],[59,22],[59,23],[59,40],[59,45],[59,46],[59,81],[59,82],[59,87],[59,88],[59,89]];
